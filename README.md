# poshman

A package manager for PowerShell modules, written in Go

## Why do this?

I use PowerShell extensively in my day-to-day work, but have always found the cmdlet structure tedious to have to remember and type in for each module that I have installed.
There is also no way to easily check for updates on the modules I have installed on my system.

## Why Go?

There are a few reasons why I decided to write this tool in Go:

1. Because I wanted to.
   - I know that sounds obnoxious, but I'm interested in Go and am learning Go, and figured that this would be a perfect opportunity to implement what I'm learning.
2. Go can compile to a single binary and can also cross-compile to multiiple operating systems and architectures.
   - This helps ensure that I can share this tool with anyone, and if I've compiled it for their system/architecture, they can use it without having to worry about having Go or any other dependencies the code has installed.
3. I eventually want to take advantage of Go's concurrency and parallelization patterns to make this process as snappy and efficient for PowerShell users as possible.

## How does it work?

The goal is to try and mimic a Unix/Linux-style package manager (similar to APT, DNF, brew, etc.) to manage PowerShell modules and their lifecycle(s) (e.g., searching, installing, updating, removing, etc.) without having to type the individual cmdlets for each module every time.
