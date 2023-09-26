package botEvents

import (
	"bubbles/bot/loaders"
	"bubbles/bot/structures"
	"bubbles/bot/toolbox"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/log"
	"github.com/dustin/go-humanize"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zcalusic/sysinfo"
)

var (
	mainEmbedColor = 0x2f3136
	noPermText     = "You need to have a role with `Manage Server` permission to use this command."
)

func ListenForCommand(e *events.ApplicationCommandInteractionCreate) {
	TRUE := true
	switch name := e.Data.CommandName(); name {
	case "host-stats":
		before, err := cpu.Get()
		if err != nil {
			DumpErrToConsole(err)
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			DumpErrToConsole(err)
			return
		}
		total := float64(after.Total - before.Total)

		memInfo, _ := mem.VirtualMemory()
		var si sysinfo.SysInfo
		si.GetSysInfo()
		if err := e.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				{
					Description: fmt.Sprintf("**OS:** %s\n**CPU:** %s", si.OS.Name, si.CPU.Model),
					Fields: []discord.EmbedField{
						{
							Name:   "CPU Usage",
							Value:  fmt.Sprintf("User: %.1f %%\nSys: %.1f %%", float64(after.User-before.User)/total*100, float64(after.System-before.System)/total*100),
							Inline: &TRUE,
						},
						{
							Name:   "Memory",
							Value:  fmt.Sprintf("Used: %v\nTotal: %v", humanize.IBytes(memInfo.Used), humanize.IBytes(memInfo.Total)),
							Inline: &TRUE,
						},
						{
							Name:   "\u200b",
							Value:  "\u200b",
							Inline: &TRUE,
						},
						{
							Name:   "Version",
							Value:  fmt.Sprintf("Disgo %v\nGo %v", disgo.Version, strings.TrimPrefix(runtime.Version(), "go")),
							Inline: &TRUE,
						},
						{
							Name:   "Goroutines",
							Value:  fmt.Sprintf("%v", runtime.NumGoroutine()),
							Inline: &TRUE,
						},
					},
					Color: mainEmbedColor,
				},
			},
		}); err != nil {
			DumpErrToConsole(err)
		}
		break
	case "invite":
		e.CreateMessage(discord.MessageCreate{
			Components: []discord.ContainerComponent{
				discord.ActionRowComponent{
					discord.ButtonComponent{
						Label: "Invite me!",
						Style: discord.ButtonStyleLink,
						URL:   fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%v&permissions=19456&scope=bot", e.Client().ApplicationID()),
					},
				},
			},
		})
		break
	case "stats":
		serverId, _ := e.SlashCommandInteractionData().OptInt("server")
		url, err := loaders.GetServer(*e.GuildID(), serverId)
		if err != nil {
			DumpErrToInteraction(e, err)
			return
		}
		req := retrieveAPIContent(url)
		var res structures.FSAPIRawData_DSS
		json.Unmarshal([]byte(req), &res)

		if res.Server.Name == "" ||
			res.Server.MapName == "" ||
			res.Server.Version == "" {
			emptyField := "····"
			res.Server.Name = "Server is offline"
			res.Server.MapName += emptyField
			res.Server.Version += emptyField
		}

		playerArray := string("")
		if res.Slots.Used < 1 {
			playerArray = "*No players online*"
		} else {
			playerArray = toolbox.GetPlayerInfo(res.Slots.Players)
		}

		if err := e.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				{
					Author: &discord.EmbedAuthor{
						Name: fmt.Sprint(res.Server.Name),
					},
					Fields: []discord.EmbedField{
						{
							Name:   "Version",
							Value:  fmt.Sprintf("%v", res.Server.Version),
							Inline: &TRUE,
						},
						{
							Name:   "Map",
							Value:  fmt.Sprintf("%v", res.Server.MapName),
							Inline: &TRUE,
						},
						{
							Name:  fmt.Sprintf("Players - (%v/%v)", res.Slots.Used, res.Slots.Capacity),
							Value: fmt.Sprintf("%v", playerArray),
						},
					},
					Color: mainEmbedColor,
				},
			},
		}); err != nil {
			DumpErrToConsole(fmt.Errorf("could not send message: %v", err.Error()))
		}
		break
	case "link":
		if !isGuildManager(e) {
			e.CreateMessage(discord.MessageCreate{
				Content: noPermText,
			})
			return
		}
		str, _ := e.SlashCommandInteractionData().OptString("server-url")
		id, _ := e.SlashCommandInteractionData().OptInt("id")
		if !strings.Contains(str, "dedicated-server-stats.xml") && !strings.Contains(str, "dedicated-server-stats.json") {
			e.CreateMessage(discord.MessageCreate{
				Content: "This is not a valid URL. Please try again.",
			})
			return
		}
		if strings.Contains(str, ".xml") {
			str = strings.Replace(str, ".xml", ".json", -1)
		}
		add, _ := loaders.AddServer(*e.GuildID(), id, str)
		if add != nil {
			log.Infof("Saved the server for %v, command performed by %v", toolbox.RESTGuild_Name(*e.GuildID(), loaders.TokenLoader("bot")), e.Member().User.Tag())
			e.CreateMessage(discord.MessageCreate{
				Content: "Saved successfully\nServer ID: `" + fmt.Sprint(id) + "`\nURL: `" + str + "`",
			})
		} else if add.UpsertedID != nil {
			log.Infof("Updated the server for %v, command performed by %v", toolbox.RESTGuild_Name(*e.GuildID(), loaders.TokenLoader("bot")), e.Member().User.Tag())
			e.CreateMessage(discord.MessageCreate{
				Content: "Updated successfully\nServer ID: `" + fmt.Sprint(id) + "`\nURL: `" + str + "`",
			})
		} else {
			DumpErrToInteraction(e, fmt.Errorf("failed to save the server"))
			return
		}
		break
	case "unlink":
		if !isGuildManager(e) {
			e.CreateMessage(discord.MessageCreate{
				Content: noPermText,
			})
			return
		}
		id, _ := e.SlashCommandInteractionData().OptInt("id")
		doc, err := loaders.DeleteServer(*e.GuildID(), id)
		if doc != nil {
			log.Infof("Deleted the server for %v, command performed by %v", toolbox.RESTGuild_Name(*e.GuildID(), loaders.TokenLoader("bot")), e.Member().User.Tag())
			e.CreateMessage(discord.MessageCreate{
				Content: "Deleted `" + fmt.Sprint(id) + "` successfully",
			})
		}
		if err != nil {
			DumpErrToInteraction(e, err)
			return
		}
		break
	case "fields":
		id, _ := e.SlashCommandInteractionData().OptInt("id")
		url, _ := loaders.GetServer(*e.GuildID(), id)
		req := retrieveAPIContent(url)
		var res structures.FSAPIRawData_DSS
		json.Unmarshal([]byte(req), &res)

		fieldId := string("")
		fieldBool := string("")
		fieldXZ := string("")
		for _, field := range res.Fields {
			fieldId += fmt.Sprintf("%v\n", field.ID)
			fieldBool += fmt.Sprintf("%v\n", field.IsOwned)
			fieldXZ += fmt.Sprintf("%v, %v\n", field.X, field.Z)
		}

		if len(res.Fields) < 1 {
			fieldId = "-"
			fieldBool = "No fields found on the server"
			fieldXZ = "-"
		} else if len(res.Fields) > 25 {
			fieldId = "-"
			fieldBool = "Too many fields to display here"
			fieldXZ = "-"
		}

		fieldOwnedLen := 0
		ownedFields, _ := e.SlashCommandInteractionData().OptBool("display-owned")
		if ownedFields {
			fieldId = string("")
			fieldBool = string("")
			fieldXZ = string("")
			for _, field := range res.Fields {
				if field.IsOwned {
					fieldOwnedLen++
					fieldId += fmt.Sprintf("%v\n", field.ID)
					fieldBool += fmt.Sprintf("%v\n", field.IsOwned)
					fieldXZ += fmt.Sprintf("%v, %v\n", field.X, field.Z)
				}
			}
		}

		titleStatus := fmt.Sprintf("There are %v fields on the server", len(res.Fields))
		if ownedFields {
			titleStatus = fmt.Sprintf("There are %v fields owned by farms on the server", fieldOwnedLen)
		}
		if err := e.CreateMessage(discord.MessageCreate{
			Embeds: []discord.Embed{
				{
					Title: titleStatus,
					Fields: []discord.EmbedField{
						{
							Name:   "#",
							Value:  fmt.Sprintf("%v", fieldId),
							Inline: &TRUE,
						},
						{
							Name:   "Owned",
							Value:  fmt.Sprintf("%v", fieldBool),
							Inline: &TRUE,
						},
						{
							Name:   "Coordinates (X,Z)",
							Value:  fmt.Sprintf("%v", fieldXZ),
							Inline: &TRUE,
						},
					},
					Footer: &discord.EmbedFooter{
						Text: "Server: " + res.Server.Name,
					},
					Color: mainEmbedColor,
				},
			},
		}); err != nil {
			DumpErrToConsole(fmt.Errorf("could not send message: %v", err.Error()))
		}
		break
	}
}

func DumpErrToConsole(err error) {
	log.Errorf("failed to send interaction response: %v", err.Error())
}

func DumpErrToInteraction(e *events.ApplicationCommandInteractionCreate, err error) {
	if err := e.CreateMessage(discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Title:       "There was an attempt...",
				Description: fmt.Sprintf("```%v```", err.Error()),
				Color:       0x560000,
			},
		},
	}); err != nil {
		DumpErrToConsole(err)
	}
}

func DumpErrToChannel(e *events.ApplicationCommandInteractionCreate, err error) {
	if _, err := e.Client().Rest().CreateMessage(e.Channel().ID(), discord.MessageCreate{
		Embeds: []discord.Embed{
			{
				Title:       "There was an attempt...",
				Description: fmt.Sprintf("```%v```", err.Error()),
				Color:       0x560000,
			},
		},
	}); err != nil {
		DumpErrToConsole(err)
	}
}

func isGuildManager(e *events.ApplicationCommandInteractionCreate) bool {
	member := e.Member()
	if member.Permissions.Has(discord.PermissionManageGuild) {
		return true
	} else {
		return false
	}
}

func retrieveAPIContent(url string) string {
	httpClient := &http.Client{Timeout: time.Second * 5}
	requ, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("failed to parse the API data: %v", err.Error())
	}
	requ.Header.Add("User-Agent", fmt.Sprintf("Bubbles/Go %v", strings.TrimPrefix(runtime.Version(), "go")))
	r, err := httpClient.Do(requ)
	if err != nil {
		log.Errorf("failed to connect to the server: %v", err.Error())
	}
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	return string(body)
}
