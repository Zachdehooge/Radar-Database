# Radar Database

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![Go Version](https://img.shields.io/badge/GO-1.22.2-yellow)
[![GO Snyk Check Master](https://github.com/Zachdehooge/Radar_Database/actions/workflows/snyk.yml/badge.svg?branch=main)](https://github.com/Zachdehooge/Radar_Database/actions/workflows/snyk.yml)

## About
* An application for downloading Level II data from the NEXRAD Level II archive
* Takes user input for the month, day, year, radar site code (KHTX - Huntsville AL for example), along with beginning and ending time frames in Zulu
* (Format for times goes as follows: HHMMSS) with all times needing to be in Zulu

## Installation

1. Download Github repo
2. Run `.exe` file
3. If smartscreen comes up, click more info -> run anyway // this warning is entirely harmless and only shows because the app is not signed

## Docker Installation
1. Make sure Docker is installed
2. Sign into Docker
3. Open a command line prompt
4. Run `docker login`
5. Run `docker pull zachdehooge/radar-database`
6. Run `docker run -it --rm -v .:/app zachdehooge/radar-database`

## Issues
* Be sure to open an issue and I will be more than happy to fix it!

## Roadmap
1. Give user the option to download all of the current days RADNEX archive without needing to prompt for time frame
2. More to come...

# TODO

- [ ] <a href="https://github.com/Zachdehooge/Radar-Database/issues/5">Issue #5</a>
- [ ] <a href="https://github.com/Zachdehooge/Radar-Database/issues/6">Issue #6</a>
- [ ] <a href="https://github.com/Zachdehooge/Radar-Database/issues/7">Issue #7</a>
- [ ] <a href="https://github.com/Zachdehooge/Radar-Database/issues/8">Issue #8</a>
- [ ] <a href="https://github.com/Zachdehooge/Radar-Database/issues/11">Issue #11</a>

### In Progress

- [ ] 

### Done ✓

- - [x] <a href="https://github.com/Zachdehooge/Radar-Database/issues/34">Issue #34</a>

<!-- Enter [x] when done with a TODO-->
