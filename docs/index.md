---
layout: home

hero:
    name: PRANAV PYQ FETCHER
    text: Amrita University Question Paper Retrieval CLI
    tagline: An ultra-fast, keyboard-driven terminal application to browse, preview, and download university exam papers.
    image:
        src: https://github.com/user-attachments/assets/44822d21-0e1d-4b3e-baa3-3450b5cc14fc
        alt: Amrita Vishwa Vidyapeetham Logo
    actions:
        - theme: brand
          text: Get Started ›
          link: /getting-started
        - theme: alt
          text: Releases & Binaries
          link: https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases
        - theme: alt
          text: GitHub Repository
          link: https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA
---

<div class="custom-home-content">

## Core Capabilities

- **⚡ Easy Mode (Fast Past 5-Year Pack Downloader)** &mdash; Select your Semester and Subject, then download all Midterms (P1, P2) and End-Semesters in 1 click, or separated by 2-year ranges (`2025–2026`, `2023–2024`, `2021–2022`, `2019–2020`).
- **Direct Branch & Course Selection** &mdash; Select directly from top branches like Cyber Security (`CYS`), Computer Science (`CSE`), Electronics (`ECE`), Artificial Intelligence (`AIE`), Electrical (`EEE`), Mechanical (`MECH`), `MCA`, and `M.Tech`, or explore the full DSpace archive.
- **Subject-First Exploration** &mdash; Pick your subject directly from an extensive curriculum catalog with real-time fuzzy search (`/`) before viewing papers.
- **Bulk & 2-Year Range Downloads** &mdash; Download past 5 years papers or specific 2-year exam packs in a single click with real-time progress indicators into organized folders.
- **Automatic Pre-Flight Network Ping** &mdash; Automatically tests socket connectivity to `dspace.amritanet.edu:8080` on launch, detecting campus Wi-Fi or VPN without any manual configuration.
- **Built-in Offline Demo Mode** &mdash; Browse courses, test fuzzy search, and explore simulated question papers offline when off-campus.
- **Native Browser PDF Preview** &mdash; Seamlessly launches question paper PDFs inside your default browser (Microsoft Edge, Chrome, Firefox, Brave) for high-fidelity viewing.
- **Organized Local Downloads** &mdash; Automatically organizes downloads into structured directories like `downloads/CYS/19CYS202/2023-2024/`.

---

## Terminal-First Architecture

Navigating university web portals right before exams shouldn't be stressful. **PRANAV PYQ FETCHER AMRITA** turns the slow, multi-click process of navigating the DSpace archive into an intuitive, keyboard-first terminal workflow.

### Network Connectivity Matrix

Because `dspace.amritanet.edu` is hosted on Amrita's private internal intranet, live fetching requires connectivity to the university network:

| Environment | Connection Mode | Network Status | Supported Capabilities |
|:---|:---|:---:|:---|
| **At Campus** | Connected to **Amrita Wi-Fi** | ● Connected | Direct live fetching, streaming & downloading from university archives. |
| **At Home with VPN** | Connected to **Amrita GlobalProtect VPN** | ● Connected | Full access to real-time question papers from anywhere in the world. |
| **At Home / Mobile Data** | Off-campus network | ○ Standalone Demo | **Offline Demo Mode** enabled. Browse sample courses, test filters, and explore all app features. |

---

### Binary Releases & Packages

Pre-compiled, ready-to-run releases are available for Windows, Linux, and macOS without requiring any Go installation:

- **Windows 64-bit:** [Download ZIP (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-windows-amd64.zip) &mdash; *Includes `run.bat` for one-click launch*
- **Linux 64-bit:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-linux-amd64.tar.gz)
- **macOS Apple Silicon:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-darwin-arm64.tar.gz)
- **macOS Intel:** [Download TAR.GZ (v2.0.0)](https://github.com/pkpranav-projects/PRANAV-PYQ-FETCHER-AMRITA/releases/download/v2.0.0/pranav-pyq-v2.0.0-darwin-amd64.tar.gz)

---

### Keybindings & Controls

| Key | Function |
|:---:|:---|
| **`↑` / `k`** | Move cursor up |
| **`↓` / `j`** | Move cursor down |
| **`Enter`** | Confirm selection |
| **`/`** | Activate real-time fuzzy search filter |
| **`Back`** | Return to the previous menu |
| **`Quit`** | Exit cleanly with animated spinner |

---

### Execution Lifecycle

```mermaid
flowchart TD
    A["pranav-pyq launch"] --> B{"Campus Network Pre-flight Check"}
    B -->|"Campus Wi-Fi / VPN"| C["Live Mode: dspace.amritanet.edu"]
    B -->|"Off-Campus"| D["Offline Demo Engine"]
    
    C --> M0["Mode Selection"]
    D --> M0

    M0 -->|"⚡ Easy Mode"| E1["1. Select Semester (Sem 1 to 8 / All)"]
    E1 --> E2["2. Select Branch (CYS, CSE, ECE, etc.)"]
    E2 --> E3["3. Select Subject"]
    E3 --> E4["4. Fast Downloader (Midterms + EndSem)"]
    E4 -->|"📥 Past 5 Years Pack"| D5["Download All 2021-2026 Papers in 1-Click"]
    E4 -->|"📦 2-Year Ranges"| D2["Download 25-26 / 23-24 / 21-22 / 19-20 Packs"]

    M0 -->|"Standard Mode"| M1["1. Course / Branch Select"]
    M1 --> M2["2. Select Subject"]
    M2 --> M3["3. Display Papers & Ranges"]
    
    M0 -->|"Browse Full Hierarchy"| H1["DSpace Communities → Semesters → Assessments → Papers"]
```

</div>
