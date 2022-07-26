/*
	Yet Another Logger (yal) is a simple and lightweight colored and leveled logger.
	It's designed to be used as a singleton type logging solution that can be
	dropped into scripts and CLIs without complex dependency injection or configuration.

	yal provides a top level singleton logger that is exposed to the user with the following levels
	Info, Debug, Warn, Error, Fatal.

	Call methods

	yal.Info("Hello World")

	yal.Infof("Hello World %s", "World")

	Set Properties on Logger

	yal.Log.Colors = false

	yal.Log.Level = yal.LevelWarn
*/
package yal
