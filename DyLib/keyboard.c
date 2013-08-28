#include <curses.h>
#include "keyboard.h"

void InitKeyboard() {
  initscr();
  noecho();
  cbreak();
  keypad(stdscr, TRUE);
  refresh();
}

int GetCharacter() {
  return getch();
}

void CloseKeyboard() {
  endwin();
}