#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <time.h>
#include <termios.h>
#include <unistd.h>
void nonblock(int state) {
    struct termios ttystate;
    tcgetattr(STDIN_FILENO, &ttystate);
    if (state==1) {
	ttystate.c_lflag &= ~ICANON;
	ttystate.c_cc[VMIN] = 1;
    } else if (state==0) {
	ttystate.c_lflag |= ICANON;
    }
    tcsetattr(STDIN_FILENO, TCSANOW, &ttystate);
}
int main(int argc, char *argv[]) {
    uint32_t score = 0, x, y, in;
    srand(time(NULL));
    printf("\033[2J");
    nonblock(1);
    while (1) {
	x = 1 + (int)( 50.0 * rand() / ( RAND_MAX + 1.0 ) );
	y = 1 + (int)( 23.0 * rand() / ( RAND_MAX + 1.0 ) );
	printf("\033[24;0H-----------------------------------------------------------------------------\nScore: %d\033[%d;%dHPress the [SPACE] key", score, y, x);
	fflush(stdout);
	in = fgetc(stdin);
	if (in != ' ') {
	    printf("\033[1;1H\033[2J\n          ____                         ___                 \n");
	    printf("         / ___| __ _ _ __ ___   ___   / _ \\__   _____ _ __ \n");
	    printf("        | |  _ / _` | '_ ` _ \\ / _ \\ | | | \\ \\ / / _ \\ '__|\n");
	    printf("        | |_| | (_| | | | | | |  __/ | |_| |\\ V /  __/ |   \n");
	    printf("         \\____|\\__,_|_| |_| |_|\\___|  \\___/  \\_/ \\___|_|   \n");
	    printf("\nYou did not hit the space key\n");
	    printf("Final Score: %d\n", score);
	    exit(0);
	}
	score++;
	printf("\033[%d;%dH                              ", y, x);
    }
}
