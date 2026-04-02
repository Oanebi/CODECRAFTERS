## Go Comments

In Go programming, **comments** are non-executable snippets of text that the compiler ignores during the build process. They serve as essential documentation, helping developers explain the logic behind their code, improve overall readability, and temporarily disable specific lines of code during debugging or testing phases.


### Key Functions of Comments
* **Code Documentation:** Explaining what a specific function or block of code is intended to do.
* **Readability:** Breaking up complex logic to make it easier for other developers (or your future self) to follow.
* **Testing:** Commenting out sections of code to test alternative implementations without deleting the original work.


### Types of Comments in Go

Go provides two primary ways to format comments, depending on the length and placement of the text:

* **Single-line Comments:** * Initiated with two forward slashes (`//`).
    * Everything from the slashes to the end of that specific line is ignored.
    * Can be placed on their own line or at the end of a line of code (inline).
* **Multi-line Comments:** * Start with `/*` and end with `*/`.
    * Useful for long explanations or blocking out entire chunks of code.
    * The compiler ignores all text contained between these two markers, regardless of how many lines it spans.