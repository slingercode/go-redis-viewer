package views

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/rivo/tview"
	"github.com/slingercode/redis-viewer/pkg"
)

var (
  app     *tview.Application
  pages   *tview.Pages
)

func finder(rclient *redis.Client) {
  // Basic window objects
  wkeys := tview.NewList()
  wkeys.SetBorder(true).SetTitle("Keys")

  wvalue := tview.NewTextView()
  wvalue.SetBorder(true).SetTitle("Value")

  // Create window layout
  flex := tview.NewFlex().
    AddItem(wkeys, 0, 1, true).
    AddItem(wvalue, 0, 1, false)

  // Get redis data
  keys := pkg.GetAllKeys(rclient)

  for i, key := range keys  {
    wkeys.AddItem(strconv.Itoa(i + 1) + ". " + key, "", 0, func () {
      wvalue.Clear()
      wvalue.SetText(key)
    })
  }

  pages = tview.NewPages().
    AddPage("name", flex, true, true)

  app.SetRoot(pages, true)
}

func SetupWindow(rclient *redis.Client) {
  app = tview.NewApplication()

  finder(rclient)

  if err := app.Run(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
