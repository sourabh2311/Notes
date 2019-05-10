# Connect VSCode To Your Remote

Hi! This tutorial aims in leveraging the power of Visual Studio Code to open and edit your files located remotely with the speed as if your files are in local machine. 

## Steps:

1. Install [Visual Studio Code Insiders edition](https://code.visualstudio.com/insiders/) (the extension which we are about to install, as of now, isn't compatible with stable edition). Note: Insiders edition is different from normal (stable) edition. To port your settings, extensions etc from normal VSCode, use [this](https://marketplace.visualstudio.com/items?itemName=Shan.code-settings-sync) extension.

2. Install the extension, "[Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)".

3. Check to see if you already have a SSH key. The public key is typically located at `~/.ssh/id_rsa.pub` on macOS / Linux, and at `%USERPROFILE%\.ssh\id_rsa.pub` on Windows. If you do not have a key, run the following command in a terminal / command prompt to generate a SSH key pair:  

   `ssh-keygen -t rsa -b 4096`

4. On macOS / Linux, run the following command in a local terminal, replacing the user and host name as appropriate.  

    `ssh-copy-id your-user-name-on-host@host-fqdn-or-ip-goes-here`

    On Windows, run the following commands in a local command prompt replacing the value of REMOTEHOST as appropriate.

    ```
    SET REMOTEHOST=your-user-name-on-host@host-fqdn-or-ip-goes-here

    scp %USERPROFILE%\.ssh\id_rsa.pub %REMOTEHOST%:~/tmp.pub
    ssh %REMOTEHOST% "mkdir -p ~/.ssh && chmod 700 ~/.ssh && cat ~/tmp.pub >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys && rm -f ~/tmp.pub"
    ```

5. Run **Remote-SSH: Connect to Host...** from the Command Palette (F1) and enter the host and your user on the host in the input box as follows: `user@hostname`

Thats it!