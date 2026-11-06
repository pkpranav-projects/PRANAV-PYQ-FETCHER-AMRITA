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

- **Direct Branch & Course Selection** &mdash; Select directly from top branches like Cyber Security (`CYS`), Computer Science (`CSE`), Electronics (`ECE`), Artificial Intelligence (`AIE`), Electrical (`EEE`), Mechanical (`MECH`), `MCA`, and `M.Tech`, or explore the full DSpace archive.
- **Subject-First Exploration** &mdash; Pick your subject directly from an extensive curriculum catalog with real-time fuzzy search (`/`) before viewing papers.
- **Bulk "Download ALL Papers"** &mdash; Download every question paper for a chosen subject or semester in a single click with real-time progress indicators into organized folders.
- **Automatic Pre-Flight Network Ping** &mdash; Automatically tests socket connectivity to `dspace.amritanet.edu:8080` on launch, detecting campus Wi-Fi or VPN without any manual configuration.
- **Built-in Offline Demo Mode** &mdash; Browse courses, test fuzzy search, and explore simulated question papers offline when off-campus.
- **Native Browser PDF Preview** &mdash; Seamlessly launches question paper PDFs inside your default browser (Microsoft Edge, Chrome, Firefox, Brave) for high-fidelity viewing.
- **Organized Local Downloads** &mdash; Automatically organizes downloads into structured directories like `downloads/CYS/19CYS202_Cryptography/`.

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
    
    C --> M1["1. Course / Branch Select (CYS, CSE, ECE, AIE, EEE, MECH, etc.)"]
    D --> M1
    
    M1 -->|"Quick Branch Flow"| M2["2. Select Subject (e.g. 19CYS202 Cryptography)"]
    M2 --> M3["3. Display Papers"]
    M3 -->|"📥 Download ALL"| BULK["Bulk Download All Papers to ./downloads/<Branch>/<Subject>/"]
    M3 -->|"Single Paper"| ACT{"Action"}
    ACT -->|"Open in Browser"| BROWSER["Default Browser Preview"]
    ACT -->|"Download PDF"| DISK["Save PDF to ./downloads/"]
    
    M1 -->|"Browse Full Hierarchy"| H1["DSpace Communities → Semesters → Assessments → Papers"]
```

</div>
