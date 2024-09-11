# fusioncatalyst official CLI tool


## Installation

### Installation Guide for `fsnctlst`

#### For macOS (via Homebrew):
You can install `fsnctlst` on macOS using Homebrew:

1. **Add the Homebrew tap:**
   ```bash
   brew tap fusioncatalyst/homebrews-osx-apps
   ```

2. **Install `fsnctlst`:**
   ```bash
   brew install fusioncatalyst/osx-apps/fsnctlst
   ```

3. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

---

#### For Windows (Manual Installation):

1. **Download the Windows binary:**
   Go to the [fsnctlst releases page](https://github.com/fusioncatalyst/cli/releases) and download the latest `fsnctlst_windows_x86_64.zip`.

2. **Extract the ZIP file:**
    - Right-click the downloaded ZIP file and extract it to a folder of your choice.

3. **Add the binary to your PATH (optional):**
    - Copy the `fsnctlst.exe` file to a folder in your system's PATH or add the folder containing `fsnctlst.exe` to the PATH.

4. **Verify the installation:**
   Open Command Prompt or PowerShell and run:
   ```powershell
   fsnctlst --version
   ```

---

#### For Linux:

You can install `fsnctlst` using native package managers or manually, depending on your Linux distribution.

##### For Debian/Ubuntu (via `.deb` package):
1. **Download and install the `.deb` package:**
   ```bash
   wget https://github.com/fusioncatalyst/cli/releases/download/<version>/fsnctlst_<version>_amd64.deb
   sudo dpkg -i fsnctlst_<version>_amd64.deb
   ```

2. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

##### For Fedora/Red Hat/CentOS (via `.rpm` package):
1. **Download and install the `.rpm` package:**
   ```bash
   wget https://github.com/fusioncatalyst/cli/releases/download/<version>/fsnctlst_<version>_x86_64.rpm
   sudo rpm -ivh fsnctlst_<version>_x86_64.rpm
   ```

2. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

##### For Arch Linux (via AUR package):
1. **Install via AUR helper (e.g., yay):**
   ```bash
   yay -S fsnctlst
   ```

2. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

##### For Alpine Linux (via `.apk` package):
1. **Download and install the `.apk` package:**
   ```bash
   wget https://github.com/fusioncatalyst/cli/releases/download/<version>/fsnctlst_<version>_x86_64.apk
   sudo apk add fsnctlst_<version>_x86_64.apk
   ```

2. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

##### For Termux (via `.deb` package):
1. **Download and install the Termux-specific `.deb` package:**
   ```bash
   wget https://github.com/fusioncatalyst/cli/releases/download/<version>/fsnctlst_<version>_termux.deb
   dpkg -i fsnctlst_<version>_termux.deb
   ```

2. **Verify the installation:**
   ```bash
   fsnctlst --version
   ```

---

### For Manual Installation on Linux:

1. **Download the binary for your OS and architecture** from the [fsnctlst releases page](https://github.com/fusioncatalyst/cli/releases).

2. **Make the binary executable**:
   ```bash
   chmod +x fsnctlst_linux_x86_64
   ```

3. **Move the binary to your PATH**:
   ```bash
   sudo mv fsnctlst_linux_x86_64 /usr/local/bin/fsnctlst
   ```

4. **Verify the installation**:
   ```bash
   fsnctlst --version
   ```

Let me know if you need additional information or adjustments for the guide!