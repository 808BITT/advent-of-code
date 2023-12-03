package day1;

import java.io.*;
import java.nio.file.*;
import java.time.*;

public class part1 {
    public static void main(String[] args) throws IOException {
        Instant start = Instant.now();

        Path path = Paths.get("day1/input.txt");
        BufferedReader reader = Files.newBufferedReader(path);

        int sum = 0;
        String line;
        while ((line = reader.readLine()) != null) {
            int first = -1, last = -1;
            for (char ch : line.toCharArray()) {
                if (Character.isDigit(ch)) {
                    if (first == -1) {
                        first = Character.digit(ch, 10);
                    }
                    last = Character.digit(ch, 10);
                }
            }
            if (first != -1 && last != -1) {
                sum += first * 10 + last;
            }
        }

        System.out.println(sum);

        Duration duration = Duration.between(start, Instant.now());
        System.out.println(duration.toMillis() + " ms");

        reader.close();
    }
}