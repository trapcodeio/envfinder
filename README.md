# EnvFinder 🔍

A powerful Go utility that locates and collects `.env` files while preserving directory structures.

## Why?
I got a new laptop and needed to back up all my `.env` files from various projects. I wanted a tool that could recursively search through directories, find all `.env` files, and copy them to a backup location while maintaining the original directory structure.
In order to be fast I consulted my code buddy Claude AI (Claude 3.7 Sonnet) to help me build this tool.

**FIST COMMIT** is the original code written by Claude AI.

## 🌟 Features

- **Recursive Search**: Thoroughly scans directories and subdirectories
- **Smart Exclusions**: Automatically skips common directories like `node_modules`, `.git`, etc.
- **Path Preservation**: Maintains original directory structure when copying files
- **Comprehensive Reporting**: Provides detailed output of found and copied files

## 🚀 Installation

```bash
# Clone the repository (if using git)
git clone https://github.com/trapcodeio/envfinder.git
cd envfinder

# Build the executable
go build -o envfinder envfinder.go
```

## 📋 Usage

```bash
# Basic usage (searches current directory, copies to ./envs)
./envfinder

# Search a specific directory
./envfinder -path /path/to/project

# Specify a custom output directory
./envfinder -output /path/to/backup/envs

# Combine options
./envfinder -path /path/to/project -output /path/to/backup/envs
```

## 🔍 How It Works

EnvFinder walks through your specified directory tree, looking for files that match `.env*` patterns (including variants like `.env.local` or `.env.production`).

When a matching file is found:
1. The tool records its location
2. Creates a mirrored directory structure in the output folder
3. Copies the file to its corresponding location
4. Preserves the original file permissions

### Example

If EnvFinder finds a file at:
```
/home/user/projects/awesome-app/backend/config/.env.production
```

It will copy it to:
```
./envs/home/user/projects/awesome-app/backend/config/.env.production
```

## 🛡️ Excluded Directories

EnvFinder automatically skips these directories to improve performance:

- `node_modules`
- `.git`
- `vendor`
- `dist`
- `build`
- `.vscode`
- `.idea`
- `__pycache__`
- `venv`
- `.env` (directories, not files)
- `bin`
- `obj`
- `target`
- `coverage`
- `logs`
- `tmp`/`temp`
- `cache`/`.cache`
- `public/assets`
- `public/dist`
- `public/build`

## 📊 Output Example

```
Found .env file: /home/user/projects/app1/backend/.env
Copied to: ./envs/home/user/projects/app1/backend/.env
Found .env file: /home/user/projects/app1/frontend/.env.local
Copied to: ./envs/home/user/projects/app1/frontend/.env.local
Found .env file: /home/user/projects/app2/api/.env.production
Copied to: ./envs/home/user/projects/app2/api/.env.production

--- Summary ---
Found 3 .env files:
1. /home/user/projects/app1/backend/.env
2. /home/user/projects/app1/frontend/.env.local
3. /home/user/projects/app2/api/.env.production

Successfully copied 3 .env files to ./envs:
1. ./envs/home/user/projects/app1/backend/.env
2. ./envs/home/user/projects/app1/frontend/.env.local
3. ./envs/home/user/projects/app2/api/.env.production
```

## ⚙️ Technical Details

- Written in Go for maximum performance and cross-platform compatibility
- Uses the standard library (`filepath.Walk`) for efficient directory traversal
- Handles file operations carefully to preserve permissions

## 🤖 Created by Claude AI

This tool was developed with the assistance of Claude AI (Claude 3.7 Sonnet), showcasing how AI can help create practical utilities for developers. The AI helped design:

- The core file discovery algorithm
- Smart directory exclusion logic
- Path preservation functionality
- The user-friendly CLI interface

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

*If you find this tool useful, consider starring the repository and sharing it with your fellow developers!*