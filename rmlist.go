// Command rmlist tries to remove directories and files given in config file, located at ~/.config/rmlist.cfg
// cfg file is a text file in which every line means one item
// if cfg file doesn't exist it is being created with initial content copied
// initial content has been inspired and copied from https://github.com/lahwaacz/Scripts/blob/master/rmshit.py
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	info := `Command rmlist tries to remove directories and files from cfg file /.config/rmlist.cfg
cfg file is a text file in which every line means one item
if cfg file doesn''t exist it is being created with initial content copied from https://github.com/lahwaacz/Scripts/blob/master/rmshit.py

Usage:

	rmlist`

	cfgIC := `/.adobe
/.macromedia
/.recently-used
/.local/share/recently-used.xbel
/Desktop
/.thumbnails
/.gconfd
/.gconf
/.local/share/gegl-0.2
/.FRD/log/app.log
/.FRD/links.txt
/.objectdb
/.gstreamer-0.10
/.pulse
/.esd_auth
/.config/enchant
/.spicec
/.dropbox-dist
/.parallel
/.dbus
/ca2
/ca2
/.distlib/
/.bazaar/
/.bzr.log
/.nv/
/.viminfo
/.npm/
/.java/
/.oracle_jre_usage/
/.jssc/
/.tox/
/.pylint.d/
/.qute_test/
/.QtWebEngineProcess/
/.qutebrowser/
/.asy/
/.cmake/
/.gnome/
/unison.log
/.texlive/
/.w3m/`

	log.SetFlags(log.Lshortfile)
	log.SetPrefix("(rmlist):: ")
	log.SetOutput(os.Stderr)

	fmt.Println(info)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	cfgp := usr.HomeDir + "/.config/"

	if _, err := os.Stat(cfgp); os.IsNotExist(err) {
		os.Mkdir(cfgp, os.ModePerm)
	}
	cfgp = cfgp + "rmlist.cfg"

	if _, err := os.Stat(cfgp); os.IsNotExist(err) {
		file, err := os.Create(cfgp)
		if err != nil {
			log.Fatal()
		}
		fmt.Fprintf(file, cfgIC)

		file.Close()
	}

	cfg, err := os.Open(cfgp)
	if err != nil {
		log.Fatal()
	}
	defer cfg.Close()

	scanner := bufio.NewScanner(cfg)
	for scanner.Scan() {
		s := usr.HomeDir + scanner.Text()
		err = os.RemoveAll(s)
		if err != nil {
			log.Printf("couldn't remove <%s> because: %v", s, err)
		} else {
			log.Println("removed succesfully (if existed) " + s)
		}

	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
