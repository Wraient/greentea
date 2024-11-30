# Greentea

This is a script to help me cheat in greentea 

- **autogreentea**: A Discord-integrated version.
- **greentea**: A simpler version without Discord integration.

# Installation

### For Linux

 ```
 curl -L -o autogreentea https://github.com/Wraient/greentea/raw/refs/heads/main/main
 curl -L -o greentea https://github.com/Wraient/greentea/raw/refs/heads/main/simple
 sudo chmod +x autogreentea greentea
 sudo mv autogreentea /usr/bin/
 sudo mv greentea /usr/bin/
 ```
     
 or

  ```
  wget -O autogreentea https://github.com/Wraient/greentea/raw/refs/heads/main/main
  wget -O greentea https://github.com/Wraient/greentea/raw/refs/heads/main/simple
  sudo chmod +x autogreentea greentea
  sudo mv autogreentea /usr/bin/
  sudo mv greentea /usr/bin/
  ```

### For Windows

1. **Download the binaries:**

    **autogreentea**:
     Open PowerShell or Command Prompt and run:

     ```powershell
     curl -L -o autogreentea.exe https://github.com/Wraient/greentea/raw/refs/heads/main/main.exe
     ```

     or use a browser to download it directly.

    **greentea**:
     ```powershell
     curl -L -o greentea.exe https://github.com/Wraient/greentea/raw/refs/heads/main/simple.exe
     ```

2. **Add the downloaded binaries to your PATH:**
   - Move the files to a directory included in your system's PATH, such as `C:\Windows\System32`, or manually add the location to the PATH environment variable.

3. **Run the programs:**
   - Open Command Prompt or PowerShell and type `autogreentea.exe` or `greentea.exe` to use the tools.

## Usage

### autogreentea (Discord-Integrated)
Set the required environment variables:

- `DISCORD_TOKEN`: Your Discord bot token.
- `GREENTEA_CHANNEL_ID`: Discord channel ID to post updates.

Run it as follows:

**Linux:**
```bash
DISCORD_TOKEN=<your_discord_token> GREENTEA_CHANNEL_ID=<your_channel_id> autogreentea
```

**Windows:**
```powershell
set DISCORD_TOKEN=<your_discord_token>
set GREENTEA_CHANNEL_ID=<your_channel_id>
autogreentea.exe
```

### greentea (Simple Version)
Run the simple version:

**Linux:**
```bash
greentea
```

**Windows:**
```powershell
greentea.exe
```

