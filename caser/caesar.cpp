#include <cstdint>
#include <cstdio>
#include <cstring>

void encrypt(char *plantext, char *ciphertext) {
    for (int i = 0; i < strlen(plantext); ++i) {
        if (plantext[i] >= 'a' && plantext[i] <= 'z')
            ciphertext[i] = ((plantext[i] - 'a' + 3) % 26) + 'A';
        else
            ciphertext[i] = plantext[i];
    }
}

void decode(char *plantext, char *ciphertext) {
    for (int i = 0; i < strlen(ciphertext); ++i) {
        if (ciphertext[i] >= 'A' && ciphertext[i] <= 'Z')
            plantext[i] = ((ciphertext[i] - 'A' - 3 + 26) % 26) + 'a';
        else
            plantext[i] = ciphertext[i];
    }
}

inline void help() {
    const char *name = "caesar";
    printf("Using `%s [-e|-d] [-i <text>|-f <filename>]`\n", name);
    printf("    -e                encrypt mode(default)\n");
    printf("    -d                decode mode\n");
    printf("    -i <text>         read from command line\n");
    printf("    -f <filename>     read string from file\n");
    printf("For example:\n");
    printf("    %s -e -i \"meet me after the toga party\"\n", name);
    printf("    %s -d -i \"PHHW PH DIWHU WKH WRJD SDUWB\"\n", name);
    printf("    %s -e -f \"plantext.txt\"\n", name);
    printf("    %s -d -f \"ciphertext.txt\"\n", name);
}

int main(int argc, char **argv) {
    const int maxn = 100;
    bool      encrypt_mode = true;
    // 0 read from stdin; 1 read from command line; 2 read from file
    char input_type = 0;
    char text[maxn], result[maxn];

    for (int i = 1; i < argc; ++i) {
        if (strcmp(argv[i], "-e") == 0) {
            encrypt_mode = true;
        } else if (strcmp(argv[i], "-d") == 0) {
            encrypt_mode = false;
        } else if (strcmp(argv[i], "-i") == 0) {
            if (i + 1 < argc) {
                input_type = 1;
                strcpy(text, argv[i + 1]);
                ++i;
                continue;
            } else {
                help();
                return 0;
            }
        } else if (strcmp(argv[i], "-f") == 0) {
            if (i + 1 < argc) {
                input_type = 2;
                FILE *f = fopen(argv[i + 1], "r");
                fscanf(f, "%[^\n]", &text);
                fclose(f);
                ++i;
                continue;
            } else {
                help();
                return 0;
            }
        } else {
            help();
            return 0;
        }
    }
    if (input_type == 0) {
        printf("Input the text for %s:",
               encrypt_mode ? "encryping" : "decoding");
        scanf("%[^\n]", &text);
    }

    if (encrypt_mode) {
        encrypt(text, result);
    } else {
        decode(result, text);
    }
    printf("%s", result);
    return 0;
}