# Available Commands

upper - 	Converts text to uppercase.
lower - 	Converts text to lowercase.
cap	 -      Capitalizes the first letter of every word.
title - 	Applies Title Case (ignores small words like 'and', 'the', 'of').
snake	-  Converts text to snake_case and removes special characters.
reverse	-  Reverses the characters within each individual word.
palindrome	- Flips the entire string backward.
countchar	- Counts every character (including symbols).
countwords	- Counts the total number of words.
countspace	- Counts only the empty spaces.
history	- Shows the last 5 successful transformations.
fly	- Displays a quick help menu.
exit   - 	Closes the program.

# LOGIC
1. The program uses an infinite for loop and bufio.NewReader to keep the session alive. This allows the user to perform multiple transformations without having to restart the program every time.
2. The code uses strings.Fields() to parse your input. This is a "robust" way to handle text because it automatically ignores extra spaces, ensuring that upper    hello works just as well as upper hello.
3. The snake function uses a strings.Builder to "clean" the string. It loops through every character and only keeps it if it is a letter or a number, ensuring the output is a valid snake_case format.
4. Uses a map[string]bool for "small words" (like of, the, an). It checks if a word exists in this map; if it does, it stays lowercase unless it is the first word of the string.