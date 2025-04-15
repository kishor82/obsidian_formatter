# obsidian_formatter

A small but handy Go utility to **convert Obsidian-style image links** (`![[image.png]]`) to **GitHub-friendly image links**, making your Markdown documentation render perfectly when pushed to GitHub.  
It also automatically **moves image files** to their respective folders, ensuring a clean and organized project structure.

---

## 🧠 Why?

Obsidian uses `![[image.png]]` to reference images locally, which doesn't work when rendering Markdown on GitHub or other platforms.  
This tool fixes that by:

- Converting all image references to proper GitHub markdown format:  
  `![image](https://github.com/user/repo/blob/main/path/to/image.png)`
- Moving image files to the directory where the markdown file resides, ensuring relative context
- Working recursively through the entire Obsidian vault

---

## 🚀 Features

- Recursively parses all Markdown files in your Obsidian directory
- Finds Obsidian image links (`![[...]]`)
- Moves images to the correct folder based on reference location
- Converts references to GitHub-compatible format using the provided GitHub root URL
- Outputs error if GitHub URL is not provided

---

## 🛠️ Installation & Build

Clone the repo and build the binary:

```bash
git clone https://github.com/kishor82/obsidian_formatter.git
cd obsidian_formatter
go build -o obsidian_formatter main.go
```

---

## ▶️ Usage

Run the tool from the **root directory of your Obsidian vault**:

```bash
./obsidian_formatter https://github.com/user/repo/blob/main/
```

### 🧾 Example

Before:

```markdown
![[image.png]]
```

After:

```markdown
![image](https://github.com/user/repo/blob/main/path/to/image.png)
```

---

## ⚠️ Error Handling

If no GitHub URL is passed, the tool will throw an error:

```
Error: GitHub root URL is required (e.g., https://github.com/user/repo/blob/main/)
```

Make sure you provide the full base GitHub URL that points to your repo's root directory for correct conversion.

---

## 📂 How It Works (Under the Hood)

1. **Recursively scans** all `.md` files starting from the current directory
2. Looks for Obsidian-style image links: `![[image.png]]`
3. **Moves the referenced image file** to the same directory as the markdown file (if it's not already there)
4. Replaces the original image reference with a GitHub-compatible link:

    ```
    ![image](<github-root-url>/relative/path/to/image.png)
    ```

5. Overwrites the markdown file with the updated content

---

## 📌 Use Cases

- Preparing your Obsidian vault for **clean publishing on GitHub**
- Ensuring all images are correctly referenced without manual editing
- Automating markdown image formatting for **tech blogs, open-source docs, or wikis**
- Great for **cleaning up image locations** and making your repo presentation-ready

---

## 👨‍💻 Developer Notes

- Works with `.png`, `.jpg`, `.jpeg`, `.gif`, etc.
- Make sure the images referenced exist in the vault before running
- You can safely commit changes after running this tool — your Markdown will look great on GitHub!
- Ideal for developers who use Obsidian for documentation and want seamless GitHub integration

---

## 🧹 TODO / Improvements (Optional)

- Add dry-run mode
- Add support for custom output directories
- Add support for excluding certain folders
- Unit tests for markdown parsing logic

---

## 🧑‍💻 Author

**Kishor** – [@kishor82](https://github.com/kishor82)

---

## 📄 License

MIT
