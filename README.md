# Proctree
View processes in tree like manner

## Usage

To run the tool, simply run the [binary](https://github.com/umegbewe/proctree/releases). No additional flags or arguments are required. The tree structure will be printed in the terminal.

### Example output on Linux
```
└─systemd───51*[{systemd}]
    └─systemd-journald───1*[{systemd-journald}]
    └─systemd-udevd───1*[{systemd-udevd}]
    └─crond───1*[{crond}]
    └─dbus-daemon───1*[{dbus-daemon}]
    └─polkitd───1*[{polkitd}]
    └─systemd-logind───1*[{systemd-logind}]
    └─NetworkManager───1*[{NetworkManager}]
    └─ModemManager───1*[{ModemManager}]
    └─cupsd───1*[{cupsd}]
    └─sddm───2*[{sddm}]
        └─Xorg───1*[{Xorg}]
        └─sddm-helper───1*[{sddm-helper}]
            └─startplasma-x11───1*[{startplasma-x11}]
                └─plasma_session───15*[{plasma_session}]
                    └─kded5───1*[{kded5}]
                    └─kwin_x11───1*[{kwin_x11}]
                    └─ksmserver───1*[{ksmserver}]
                    └─org_kde_powerdevil───1*[{org_kde_powerdevil}]
                    └─kaccess───1*[{kaccess}]
                    └─polkit-kde-authentication-agent-1───1*[{polkit-kde-authentication-agent-1}]
                    └─xembedsniproxy───1*[{xembedsniproxy}]
                    └─baloo_file───1*[{baloo_file}]
                    └─plasmashell───6*[{plasmashell}]
                        └─brave-browser───3*[{brave-browser}]
.....
```
## Features

* Display the parent-child relationships of all processes running on the system
* Indicate the number of child processes for each parent process
* Use a tree structure to clearly show the relationships between processes

## Contribution

Feel free to contribute to this project by submitting pull requests or opening issues.

## License
This project is licensed under the MIT License. See the [LICENSE](https://github.com/umegbewe/proctree/blob/main/LICENSE) file for details.
