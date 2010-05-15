#!/usr/bin/env python
import urwid
import time
txt = urwid.Text("push the spacebar!")
def show_or_exit(input):
	if input != ' ':
		txt.set_text("You didn't press space!")
		time.sleep(2) 
		raise urwid.ExitMainLoop()
	txt.set_text("you pushed the spacebar!")

fill = urwid.Filler(txt, 'top')
loop = urwid.MainLoop(fill, unhandled_input=show_or_exit)
loop.run()