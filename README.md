# PRANAV PYQ FETCHER AMRITA 🚀

![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Version](https://img.shields.io/badge/version-v2.0.0-emerald)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-blue)

An enhanced, interactive Command-Line Interface (TUI) tool to browse, preview, and download **Previous Year Question Papers (PYQs)** from Amrita Vishwa Vidyapeetham's internal DSpace library repository.

---

## ✨ Features

- **🎓 Full Academic Hierarchy Navigation**: Browse Course &rarr; Semester &rarr; Assessment (Periodical 1, Periodical 2, End Sem) &rarr; Subject &rarr; Question Paper.
- **⚡ Pre-flight Campus Network Detection**: Automatically detects if you are connected to the Amrita campus intranet.
- **🚀 Offline Demo Mode**: Test and explore all menus, subjects, and simulated question papers even when you're at home off-campus!
- **🌐 In-Browser PDF Viewing**: Opens question papers directly in your default web browser (Edge, Chrome, Firefox, etc.).
- **📥 Clean Downloads Folder**: Automatically saves question paper PDFs into a dedicated `downloads/` folder.
- **🔍 Quick Fuzzy Filter**: Press `/` in any selection menu to quickly filter through dozens of question papers or subjects.
- **⬅️ Interactive Breadcrumbs & Back Navigation**: Keep track of your path (`Selection: B.Tech CSE → Sem 4 → End Sem`) and easily backtrack without restarting.

---

## 📶 Network Requirements

The Amrita University digital question paper repository (`dspace.amritanet.edu:8080`) is hosted on the internal university intranet:

| Environment | Connectivity Status | App Capability |
|---|---|---|
| **At College (Campus Wi-Fi)** | Connected to `Amrita` Wi-Fi | **Full Live Fetching** & real-time downloads from DSpace |
| **At Home / Off-Campus** | Personal Wi-Fi / Mobile Data | **Offline Demo Mode** (explore menus & preview flow) |
| **At Home with VPN** | Connected to Amrita GlobalProtect VPN | **Full Live Fetching** & real-time downloads from DSpace |

> 💡 **Tip:** You do **not** need to worry if you are running this from home right now — the app will gracefully detect that you're off-campus and offer you the **Offline Demo Mode** so you can test it immediately!

---

## 🚀 Quick Start & Usage

### Option 1: Direct Executable (Recommended)
You can directly run the pre-compiled binary:
- **Windows (Double-Click):** Double-click [`run.bat`](file:///run.bat) in this folder.
- **Or via Terminal:**
  ```powershell
  .\pranav-pyq.exe
  ```

### Option 2: Run with Go
```bash
go run main.go
```

### Checking the Version
```bash
.\pranav-pyq.exe version
```

---

## ⌨️ Navigation Controls

- **`↑` / `↓` Arrow Keys**: Navigate through courses, semesters, and question papers.
- **`Enter`**: Confirm selection.
- **`/` Key**: Open search filter to type and narrow down options.
- **`Back`**: Return to the previous menu.
- **`Quit`**: Exit the application cleanly.

---

## 🛠️ Built With

- **Language:** Go (Golang 1.23+)
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **TUI Prompts & Forms:** [Charm Huh](https://github.com/charmbracelet/huh)
- **Styling & Layout:** [Charm Lipgloss](https://github.com/charmbracelet/lipgloss)
- **HTML Parsing & Web Scraping:** [Soup](https://github.com/anaskhan96/soup)
