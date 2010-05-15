#!/usr/bin/env python
# 2010. GPL v3 and all later versions. Jrabbit.
import urwid
txt = urwid.Text("push the spacebar!")
#Should I even add physics?
physics_fact=["f=ma","Vx=-9.8m/s^2 + V0*COSTheta"]
#reddit meme goes here
def show_or_exit(input):
	if input != ' ':
		# You didn't press space!
		raise urwid.ExitMainLoop()
	txt.set_text("you pushed the spacebar!")

fill = urwid.Filler(txt, 'top')
loop = urwid.MainLoop(fill, unhandled_input=show_or_exit)
loop.run()
print "You didn't press space!"
