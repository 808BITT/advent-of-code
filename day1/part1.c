#include <stdio.h>
#include <time.h>
#include <windows.h>

int main()
{
    FILE *file = fopen("input.txt", "r");
    if (file == NULL)
    {
        printf("Could not open file\n");
        return 1;
    }

    LARGE_INTEGER frequency;
    LARGE_INTEGER start, end;
    QueryPerformanceFrequency(&frequency);
    QueryPerformanceCounter(&start);

    int sum = 0;
    char line[256];
    while (fgets(line, sizeof(line), file))
    {
        int first = -1, last = -1;
        for (int i = 0; line[i] != '\0'; i++)
        {
            if (line[i] >= '0' && line[i] <= '9')
            {
                if (first == -1)
                {
                    first = line[i] - '0';
                }
                last = line[i] - '0';
            }
        }
        if (first != -1 && last != -1)
        {
            sum += first * 10 + last;
        }
    }

    printf("%d\n", sum);

    QueryPerformanceCounter(&end);
    long long microseconds = (end.QuadPart - start.QuadPart) * 1000000 / frequency.QuadPart;

    printf("%lld μs\n", microseconds);

    fclose(file);
    return 0;
}