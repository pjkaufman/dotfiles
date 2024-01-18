# Magnum

Magnum is a program that checks if the list of specified light novels has any updates and notes the release dates of any new entries.

## Supported Publishers

- JNovel Club
- Yen Press

## Light Novels to Account for

Daily Life of the Immortal King - Novel Updates?
Eighth Son - Novel Updates
Moonlit Fantasy - Novel Updates?
The rising of the shield hero - Open Peace Books?  https://www.onepeacebooks.com/jt/ShieldHeroLNV.html

## Features to Add

- Show Info should display light novels releasing this month in yellow, and last month or earlier in red
- Get-info should be able to determine if the latest release for a series already came out and handle that accordingly
- Add robots.txt checks to http calls to sites like Wikipedia and JNovels
  - This should respect whether or not magnum is allowed to read from those sites
  - The robots.txt should be cached per run (something like storing it in a map of site name to robots.txt content)
- Add more unit tests and validation for commands and parsing logic to make sure it works as intended and is easier to refactor down the road since breaking changes should be easier to catch
