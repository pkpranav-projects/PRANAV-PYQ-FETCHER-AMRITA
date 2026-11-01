---
layout: home

hero:
    name: PRANAV PYQ FETCHER
    text: Next-Gen Amrita PYQs Retrieval
    tagline: An ultra-fast, keyboard-driven Terminal CLI tool to browse, preview, and download university exam papers.
    image:
        src: https://github.com/user-attachments/assets/44822d21-0e1d-4b3e-baa3-3450b5cc14fc
        alt: Amrita Vishwa Vidyapeetham Logo
    actions:
        - theme: brand
          text: 🚀 Quick Start
          link: /getting-started
        - theme: alt
          text: 📦 Releases & Downloads
          link: https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases
        - theme: alt
          text: 💻 View GitHub
          link: https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA

features:
    - icon: ⚡
      title: Automatic Pre-Flight Network Ping
      details: Pings the internal Amrita repository on launch to detect campus Wi-Fi or VPN connectivity automatically.
    - icon: 🚀
      title: Built-in Offline Demo Mode
      details: Stuck at home? Browse sample courses, test fuzzy search, and explore simulated question papers offline.
    - icon: 🔍
      title: Instant Search Filter (/)
      details: Tap / in any menu to type and filter subjects, semesters, and question papers in milliseconds.
    - icon: 🌐
      title: Native Browser PDF Preview
      details: Seamlessly launches question paper PDFs inside Microsoft Edge, Chrome, Firefox, or Brave.
    - icon: 📥
      title: Organized Local Downloads
      details: Automatically creates and routes PDF downloads into a clean downloads/ folder.
    - icon: ⬅️
      title: Interactive Stack Breadcrumbs
      details: Keep track of your exact path (e.g. B.Tech CSE → Sem 4 → End Sem) and step backwards effortlessly.
---

<div class="custom-home-content">

## 🌟 The Ultimate Exam Prep Companion

Navigating university web portals right before exams shouldn't be stressful. **PRANAV PYQ FETCHER AMRITA** turns the slow, multi-click process of navigating the DSpace archive into an intuitive, keyboard-first terminal workflow.

### 📶 Network Connectivity Matrix

Because `dspace.amritanet.edu` is hosted on Amrita's private internal intranet, live fetching requires connectivity to the university network:

| Environment | Connection Mode | Status | Supported Features |
|:---|:---|:---:|:---|
| **At Campus** | Connected to **Amrita Wi-Fi** | 🟢 Live | Direct live fetching, streaming & downloading from university archives. |
| **At Home with VPN** | Connected to **Amrita GlobalProtect VPN** | 🟢 Live | Full access to real-time question papers from anywhere in the world. |
| **At Home / Mobile Data** | Off-campus network | 🟣 Demo Mode | **Offline Demo Mode** enabled! Browse sample courses, test filters, and test all app features. |

---

### 📦 Quick Downloads

Pre-compiled, ready-to-run releases are available for Windows, Linux, and macOS without requiring any Go installation:

- **Windows 64-bit:** [Download ZIP (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-windows-amd64.zip) &mdash; *Includes `run.bat` for one-click launch*
- **Linux 64-bit:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-linux-amd64.tar.gz)
- **macOS Apple Silicon:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-darwin-arm64.tar.gz)
- **macOS Intel:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-darwin-amd64.tar.gz)

---

### ⌨️ Navigation Cheatsheet

| Key | Function |
|:---:|:---|
| **`↑` / `k`** | Move cursor up |
| **`↓` / `j`** | Move cursor down |
| **`Enter`** | Confirm selection |
| **`/`** | Activate real-time fuzzy search filter |
| **`Back`** | Return to the previous menu |
| **`Quit`** | Exit cleanly with animated spinner |

---

### 🏗️ How It Works

```mermaid
flowchart TD
    A["🚀 Launch: pranav-pyq"] --> B{"Campus Network Pre-flight Check"}
    B -->|"Campus Wi-Fi / VPN"| C["🟢 Live Mode (dspace.amritanet.edu)"]
    B -->|"Off-Campus"| D["🟣 Offline Demo Mode (Mock Engine)"]
    
    C --> M1["1. Course Selection (B.Tech, MCA, M.Tech, etc.)"]
    D --> M1
    
    M1 --> M2["2. Semester Selection (Sem 1 to Sem 8)"]
    M2 --> M3["3. Assessment Type (P1, P2, End Sem)"]
    M3 --> M4["4. Subject / Sub-Community"]
    M4 --> M5["5. Question Paper Year & File"]
    
    M5 --> ACT{"User Action"}
    ACT -->|"🌐 Open in Browser"| BROWSER["Executes OS Protocol (rundll32 / open / xdg-open)"]
    ACT -->|"📥 Download PDF"| DISK["Streams file into ./downloads/ folder"]
```

</div>
