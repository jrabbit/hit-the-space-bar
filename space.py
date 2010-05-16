#!/usr/bin/env python
# 2010. GPL v3 and all later versions. Jrabbit.
import urwid
txt = urwid.Text("push the spacebar!")
#Should I even add physics?
physics_fact=["f=ma","Vx=-9.8m/s^2 + V0*COSTheta"]
score = 1
#reddit meme goes here
def show_or_exit(input):
	global score
	if input != ' ':
		# You didn't press space!
		raise urwid.ExitMainLoop()
	if score < 2:
		message = "you pushed the spacebar!"
	else:
		message = "you pushed the spacebar " + str(score) + " times!"
	txt.set_text(message)
	score = score + 1
fill = urwid.Filler(txt, 'top')
loop = urwid.MainLoop(fill, unhandled_input=show_or_exit)
loop.run()
print "You didn't press space!"
#YOU'RE THE NIGHT SURGEON!