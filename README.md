# Radar Database

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![Go Version](https://img.shields.io/badge/GO-1.22.2-yellow)
[![GO Snyk Check](https://github.com/Zachdehooge/Radar_Database/actions/workflows/snyk.yml/badge.svg)](https://github.com/Zachdehooge/Radar_Database/actions/workflows/snyk.yml)

# About
* An application that takes user input for the month, day, year, radar site code (KHTX - Huntsville AL for example), beginning and ending time frame that you would like radar data downloaded from the NWS level II radar archive
* (Format goes as follows: HHMMSS) with all times needing to be in Zulu
* *NOTE: Be sure to make the time at the top of the hour 000100 for example of 1am and the end time can be no less than two hours after the start time, I am working to fix this issue so a start time of 000100 would need a end time no less than 000300 or 3am*

# Installation

1. Download Github repo
2. Run `.exe` file

# Issues
* Be sure to open an issue and I will be more than happy to fix it!

# Roadmap
1. Give user the option to download all of the current days RADNEX archive without needing to prompt for time frame
2. More to come...
