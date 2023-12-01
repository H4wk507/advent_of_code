#include <stdio.h>
#include <stdlib.h>

int get_filesize(const char *filename) {
    FILE *fptr = fopen(filename, "r");
    fseek(fptr, 0L, SEEK_END);
    int size = ftell(fptr);
    fseek(fptr, 0L, SEEK_SET);
    fclose(fptr);
    return size;
}

char *read_file(const char *filename, int size) {
    FILE *fptr = fopen(filename, "r");
    char *data = (char *)malloc(size);
    fread(data, 1, size, fptr);
    fclose(fptr);
    return data;
}

void part1(const char *data, int size) {
    int floor = 0;
    for (int i = 0; i < size; i++) {
        switch (data[i]) {
            case '(': floor++; break;
            case ')': floor--; break;
        }
    }
    printf("%d\n", floor);
}

void part2(const char *data, int size) {
    int floor = 0;
    int position = -1;
    for (int i = 0; i < size; i++) {
        switch (data[i]) {
            case '(': floor++; break;
            case ')': floor--; break;
        }
        if (floor == -1) {
            position = i + 1;
            break;
        }
    }
    printf("%d\n", position);
}

int main(void) {
    int size = get_filesize("data.txt");
    char *data = read_file("data.txt", size);
    part1(data, size);
    part2(data, size);
    free(data);
}
