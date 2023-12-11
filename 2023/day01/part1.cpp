#include <iostream>
#include <fstream>
#include <chrono>
#include <istream>
#include <string>

int main()
{
    std::ifstream file("input.txt");
    if (!file.is_open())
    {
        std::cout << "Could not open file\n";
        return 1;
    }

    auto start = std::chrono::high_resolution_clock::now();

    int sum = 0;
    std::string line;
    while (std::getline(file, line))
    {
        int first = -1, last = -1;
        for (char ch : line)
        {
            if (std::isdigit(ch))
            {
                if (first == -1)
                {
                    first = ch - '0';
                }
                last = ch - '0';
            }
        }
        if (first != -1 && last != -1)
        {
            sum += first * 10 + last;
        }
    }

    std::cout << sum << "\n";

    auto end = std::chrono::high_resolution_clock::now();
    auto duration = std::chrono::duration_cast<std::chrono::microseconds>(end - start);
    std::cout << duration.count() << " μs\n";

    file.close();
    return 0;
}