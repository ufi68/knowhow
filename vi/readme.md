# vi / vim

## Aufruf 
```
VIM - Vi IMproved 8.0 (2016 Sep 12, compiled Jun 14 2022 02:52:30)

Verwendung: vim [Argumente] [Datei ..]      editiere die angegebenen Datei(-en)
   oder: vim [Argumente] -               lese Text von stdin
   oder: vim [Argumente] -t tag          öffne Datei in der der Tag definiert wurde
   oder: vim [Argumente] -q [Fehler-Datei]  öffne Datei mit erstem Fehler

Argumente:
   --                   Hiernach nur Dateinamen
   -v                   Vi Modus (wie "vi")
   -e                   Ex Modus (wie "ex")
   -E                   Verbesserter Ex Modus
   -s                   Leiser (batch) Modus (nur für "ex")
   -d                   Diff Modus (wie "vimdiff")
   -y                   Leichter Modus (wie "evim", ohne Modi)
   -R                   Nur Lese-Modus (wie "view")
   -Z                   Eingeschränkter Modus (wie "rvim")
   -m                   Modifikationen (Schreiben von Dateien) sind nicht erlaubt
   -M                   Modifikationen im Text nicht erlaubt
   -b                   Binärmodus
   -l                   Lisp Modus
   -C                   Kompatibel zu Vi: 'compatible'
   -N                   Nicht ganz kompatibel zu Vi: 'nocompatible'
   -V[N][Datei]         Verbose [level N] [Logge nach Datei]
   -D                   Debug-Modus
   -n                   Keine Auslagerungsdatei, verwende nur Speicher
   -r                   Liste nur Auslagerungsdateien auf
   -r (mit Dateiname)   Stelle abgestürzte Session wieder her
   -L                   Genauso wie -r
   -A                   Start im Arabischen Modus
   -H                   Start im Hebräischen Modus
   -F                   Start im Farsi Modus
   -T <terminal>        Setze Terminaltyp auf <terminal>
   --not-a-term         Keine Warnung ausgeben, wenn Eingabe/Ausgabe nicht auf einem Terminal ausgegebn wird
   --ttyfail            Beende, wenn Ein- oder Ausgabe nicht auf einem Terminal ausgegeben wird
   -u <vimrc>           Benutze <vimrc> anstatt jeglicher .vimrc
   --noplugin           lade keine "plugin"-Skripte
   -p[N]                Öffne N Tabs (Vorgabe: einen für jede Datei)
   -o[N]                Öffne N Fenster (Vorgabe: für jede Datei eins)
   -O[N]                Wie -o, aber teile vertikal
   +                    Starte am Ende der Datei
   +<lnum>              Start in Zeile <lnum>
   --cmd <Befehl>       Führe <Befehl> vor dem Laden jeglicher vimrc-Datei aus
   -c <Befehl>          Führe <Befehl> nach dem Laden der ersten Datei aus
   -S <session>         Lese Datei <session> nach dem Laden der ersten Datei
   -s <scriptin>        Lese Normal-Modus Befehle aus der Skript-Datei <scriptin>
   -w <scriptout>       Alle getippten Befehle der Skript-Datei <scriptout> anfügen
   -W <scriptout>       Schreibe getippte Befehle in die Skript-Datei <scriptout>
   -x                   Editiere verschlüsselte Dateien
   --startuptime <Datei>        Schreibe Start Zeitmessung in <Datei>
   -i <viminfo>         Benutze <viminfo> statt .viminfo
   --clean              'nocompatible', Vim Standardwerte, keine Plugins, keine Viminfo
   -h  or  --help       Anzeigen der Hilfe (diesen Text) und beenden
   --version            Versionsinformation anzeigen und beenden
   ```
   
## Modi 
   
### Kommandozeilen-Modus 
```   
:w - Speichern 
:q - Beenden 
```
   
### Befehlsmodus 
  
### Einfügemodus 
   
   
