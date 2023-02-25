package main

import (
	"fmt"
	"github.com/adrg/xdg"
	"github.com/cmuench/krunner-gitlab/pkg/runner"
	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

// see for definition /usr/share/dbus-1/interfaces/kf5_org.kde.krunner1.xml
const intro = `
<node>
  <interface name="org.kde.krunner1">
    <method name="Actions">
      <annotation name="org.qtproject.QtDBus.QtTypeName.Out0" value="RemoteActions" />
      <arg name="matches" type="a(sss)" direction="out" />
    </method>
    <method name="Run">
      <arg name="matchId" type="s" direction="in"/>
      <arg name="actionId" type="s" direction="in"/>
    </method>
    <method name="Match">
      <arg name="query" type="s" direction="in"/>
      <annotation name="org.qtproject.QtDBus.QtTypeName.Out0" value="RemoteMatches"/>
      <arg name="matches" type="a(sssuda{sv})" direction="out"/>
    </method>
  </interface>` + introspect.IntrospectDataString + `</node>`

func main() {
	// read config
	viper.SetConfigName("config")
	viper.SetDefault("items_to_show", 10)
	viper.SetDefault("query_prefix", "gitlab")
	viper.SetDefault("query_min_length", 4)
	viper.AddConfigPath(xdg.ConfigHome + "/krunner-gitlab")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// configure GitLab client connection
	client, err := gitlab.NewClient(viper.GetString("token"), gitlab.WithBaseURL(viper.GetString("url")))
	if err != nil {
		panic(err)
	}

	// connect to Session DBUS
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	reply, err := conn.RequestName("de.cmuench.gitlab", dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		panic("Name de.cmuench.gitlab already taken")
	}

	// create & export runner instance
	f := runner.NewRunner(
		client,
		viper.GetString("query_prefix")+" ",
		viper.GetInt("query_min_length"),
	)

	err = conn.Export(f, "/krunner", "org.kde.krunner1")
	if err != nil {
		panic(err)
	}

	err = conn.Export(introspect.Introspectable(intro), "/krunner", "org.freedesktop.DBus.Introspectable")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on de.cmuench.gitlab/krunner...")
	select {}
}
