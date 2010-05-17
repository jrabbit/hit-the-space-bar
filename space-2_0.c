#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <time.h>
#include <termios.h>
#include <unistd.h>
void nonblock(int state) {
    struct termios ttystate;
    tcgetattr(STDIN_FILENO, &ttystate);             //            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
    if (state==1) {                                 //                    Version 2, December 2004
	ttystate.c_lflag &= ~ICANON;                //
	ttystate.c_cc[VMIN] = 1;                    // Copyright (C) 2010 Eric Kilfoil <ekilfoil@gmail.com>
    } else if (state==0) {                          //
	ttystate.c_lflag |= ICANON;                 // Everyone is permitted to copy and distribute verbatim or modified
    }                                               // copies of this license document, and changing it is allowed as long
    tcsetattr(STDIN_FILENO, TCSANOW, &ttystate);    // as the name is changed.
}                                                   //
int main(int argc, char *argv[]) {                  //            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
    uint32_t score = 0, x, y, in, vikings = 0;      //   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
    srand(time(NULL));                              //
    printf("\033[2J");                              //  0. You just DO WHAT THE FUCK YOU WANT TO.
    nonblock(1);
    while (1) {
	x = 1 + (int)( 50.0 * rand() / ( RAND_MAX + 1.0 ) );
	y = 1 + (int)( 23.0 * rand() / ( RAND_MAX + 1.0 ) );
	printf("\033[24;0H-----------------------------------------------------------------------------\nScore: %d         Vikings: %d\033[%d;%dHPress the [SPACE] key", score, vikings, y, x);
	fflush(stdout);
	in = fgetc(stdin); vikings += rand() < RAND_MAX / 4 ? 1 : 0;
	if (in != ' ') {
	    printf("\033[1;1H\033[2J\n          ____                         ___                 \n");
	    printf("         / ___| __ _ _ __ ___   ___   / _ \\__   _____ _ __ \n");
	    printf("        | |  _ / _` | '_ ` _ \\ / _ \\ | | | \\ \\ / / _ \\ '__|\n");
	    printf("        | |_| | (_| | | | | | |  __/ | |_| |\\ V /  __/ |   \n");
	    printf("         \\____|\\__,_|_| |_| |_|\\___|  \\___/  \\_/ \\___|_|   \n");
	    printf("\nYou did not hit the space key\n");
	    printf("Final Score: %d (Vikings: %d)\n", score, vikings);
	    exit(0);
	}
	score++;
	printf("\033[%d;%dH                              ", y, x);
    }
}
