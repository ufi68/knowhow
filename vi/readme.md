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

Wechsel mit `<ESC>`
   
### Kommandozeilen-Modus 
```   
:w - Speichern 
:w <filename> - Speichern als Datei <filename>
:q - Beenden 
:q! - Beenden ohne Änderungen zu speichern 
```
   
### Befehlsmodus 
https://www-user.tu-chemnitz.de/~hot/VIM/vi.html

#### Text eingeben
```
a - Text nach dem Cursor einfügen
A - Text am Ende der Zeile einfügen
i - Text vor dem Cursor einfügen
I - Text am Anfang der Zeile einfügen
o - unter dem Cursor Leerzeile für Texteingabe eröffnen
O - über dem Cursor Leerzeile für Texteingabe eröffnen
:r file - file lesen und nach der aktuellen Zeile einfügen
:nr file - file lesen und nach der n-ten Zeile einfügen
``` 

#### Rollen und Ausrichten des Bildschirminhaltes
```
<CTRL>-f      eine Seite vorwärts rollen
<CTRL>-b      eine Seite rückwärts rollen
<CTRL>-d      eine halbe Seite vorwärts rollen
<CTRL>-u      eine halbe Seite rückwärts rollen

z <RET>       aktuelle Zeile wird oberste Zeile
              des Bildschirms
z.            aktuelle Zeile wird zur mittlere Zeile
z-            aktuelle Zeile wird zur untersten Zeile
```  

#### Bewegen des Cursors in einem File
```  
h oder <--     Cursor nach links bewegen
l oder -->     Cursor nach rechts bewegen
j              Cursor nach unten bewegen
k              Cursor nach oben bewegen
$              Cursor zum Ende der aktuellen Zeile bewegen
0 (Null)       Cursor zum Anfang der aktuellen Zeile bewegen
^              Cursor zum ersten sichtbaren Zeichen der aktuellen Zeile bewegen
tchar          Cursor vor erstes Zeichen char bewegen
ntchar         Cursor vor n-tes Zeichen char bewegen

w              Cursor zum Anfang des nächsten Wortes bewegen
W              Cursor zum Anfang des nächsten Wortes bewegen
               (ohne Berücksichtigung von Sonderzeichen)
b              Cursor zum Anfang des vorhergehenden Wortes bewegen
B              Cursor zum Anfang des vorhergehenden Wortes bewegen (ohne Berücksichtigung von Sonderzeichen)
e              Cursor zum Ende des nächsten Wortes bewegen
E              Cursor zum Ende des nächsten Wortes bewegen
               (ohne Berücksichtigung von Sonderzeichen)
 
H              Cursor auf oberste Zeile des Bildschirms
M              Cursor auf mittlere Zeile des Bildschirms
L              Cursor auf unterste Zeile des Bildschirms
1G oder gg     Cursor auf 1. Zeile im File bewegen
nG             Cursor auf n-te Zeile im File bewegen
G              Cursor auf letzte Zeile im File bewegen
n+             Cursor n Zeilen vorwärts bewegen
n-             Cursor n Zeilen rückwärts bewegen

(              Cursor zum Anfang des Satzes bewegen
)              Cursor zum Anfang des nächsten Satzes bewegen
{              Cursor zum Anfang des Absatzes bewegen
}              Cursor zum Anfang des nächsten Absatzes bewegen
% (auf {}[]()) Cursor auf zugehörende Klammer bewegen
```  

#### Text löschen
Im vi-Kommandomodus:
```
x             durch Cursor markiertes Zeichen löschen
nx            n Zeichen ab Cursor löschen
dd            aktuelle Zeile löschen
ndd           die nächsten n Zeilen löschen (einschl. der aktuellen)
dw            aktuelles Wort löschen (vom Cursor bis Wortende)
dW            aktuelles Wort einschl. Sonderzeichen löschen
ndw           n Wörter ab markiertem Wort löschen
db            vorhergehendes Wort löschen
dB            vorhergehendes Wort einschl. Sonderzeichen löschen
ndb           n Wörter vor dem Cursor löschen
d$ oder D     aktuelle Zeile vom Cursor bis Zeilenende löschen
d0 (Null)     aktuelle Zeile vom Cursor bis Zeilenanfang löschen
dcursor_cmd   Text vom Cursor bis zu der durch cursor_cmd bestimmten Stelle löschen
:m,nd         Zeilen m bis n löschen
```

Im vi-Eingabemodus:
```
<CTRL>-h oder Backspace  vorhergehendes Zeichen löschen
<CTRL>-w                 vorhergehendes Wort löschen
<CTRL>-u                 alle eingebenen Zeichen der aktuellen Zeile komplett löschen
```

#### Ändern von Text
```
rchar              aktuelles Zeichen durch char ersetzen
Rtext              aktuelle(s) Zeichen durch text ersetzen
stext              aktuelles Zeichen durch text ersetzen
Stext oder cctext  aktuelle Zeile durch text ersetzen
cwtext             aktuelles Wort durch text ersetzen
ncwtext            n Wörter ab Cursor durch text ersetzen
Ctext              aktuelle Zeile von Cursor bis Zeilenende durch text ersetzen
ccursor_cmd text   Text von Cursor bis cursor_cmd durch text ersetzen
~                  aktueller Buchstabe wird in Groß- bzw. Kleinbuchstabe umgewandelt
n~                 n Buchstaben ab Cursor in Groß- bzw. Kleinbuchstaben umwandeln
.                  letztes Kommando wiederholen
```

#### Änderungen zurücksetzen
```
u             zuletzt ausgeführten Befehl rückgängig machen
U             Zustand der aktuellen Zeile vor der Änderung herstellen
:q!           vi-Sitzung ohne Sicherung der Änderung verlassen
```

#### Text kopieren oder verschieben
```
Y oder yy     aktuelle Zeile zwischenspeichern
nY oder nyy   n Zeilen ab Cursor zwischenspeichern
ycursor_cmd   Text von Cursor bis cursor_cmd zwischenspeichern

dd            aktuelle Zeile löschen
ndd           n Zeilen ab Cursor löschen
dcursor_cmd   Text von Cursor bis cursor_cmd löschen

p             gelöschten oder zwischengespeicherten Text hinter (oder unter) dem Cursor einfügen
P             gelöschten oder zwischengespeicherten Text vor (oder über) dem Cursor einfügen

xp            zwei Zeichen vertauschen

J             zwei Zeilen zusammenfügen
nJ            n Zeilen zusammenfügen
```

#### Globales Suchen und Ersetzen von Text
```
:m,ns/str1/str2/    von Zeile m bis Zeile n wird str1 durch str2 ersetzt
:m,ns/str1/str2/g   von Zeile m bis Zeile n werden alle str1 durch str2 ersetzt
&                   das letzte :s-Kommando wird wiederholt

:g/str/cmd          cmd wird für alle Zeilen ausgeführt, die str enthalten
:v/str/cmd          cmd wird für alle Zeilen ausgeführt, die str nicht enthalten
```

#### Suchen im Text
```
fchar         vorwärts suchen von char in der aktuellen Zeile
Fchar         rückwärts suchen von char in der aktuellen Zeile
/str<RET>     vorwärts suchen von str ab der aktuellen Zeile
?str<RET>     rückwärts suchen von str ab der aktuellen Zeile
n             nächstes Muster (vorwärts) suchen
```

## Mehrere Fenster

#### Ccommand line 
```
$ vim -o file1.txt file2.txt
```
#### Split Windows inside vi
All vim windows related commands begin with `<Ctrl>-w`
```
<Ctrl>-w s          Split Horizontal 
<Ctrl>-w <Ctrl>-w   Jump to next Window (toggle)
<Ctrl>-w <Ctrl>-j   Jump to Window down
<Ctrl>-w <Ctrl>-k   Jump to Window up
<Ctrl>-w <Ctrl>-l   Jump to Window right
<Ctrl>-w <Ctrl>-w   Jump to Window left

:vsplit             Split vertical
```

#### Falten 
```
[count]zf           Erstellen einer Faltung[count]
zC                  Schließen einer Faltung
zO                  Öffnen einer Faltung 
zD                  Löschen einer Faltung
```



####  Shellkommandos
```
:!cmd         Kommando cmd wird der Shell zur Ausführung übergeben
:r !cmd       Kommando cmd wird ausgeführt und Ergebnis hinter der aktuellen Zeile eingefügt
```

#### Vi-Optionen
Durch das Belegen von vi-Optionen (Kommando :set) mit einem Wert oder das Setzen von Optionen kann die Arbeitsweise des Editors beeinflusst werden. Die Kommandos set können in ein File mit dem Namen $HOME/.exrc eingetragen werden, das bei jedem vi-Editoraufruf abgearbeitet wird.
```
:set all       Anzeige Belegung aller Optionen

:set number          Zeilennummern anzeigen
:set nonumber        Zeilennummern nicht anzeigen
:set ignorecase      beim Suchen soll nicht zwischen Groß- und Kleinbuchstaben unterschieden werden
:set noignorecase    beim Suchen soll zwischen Groß- und Kleinbuchstaben unterschieden werden
:set showmatch bei   Eingabe einer schließenden Klammer wird die dazugehörige öffnende Klammer angezeigt
:set noshowmatch     keine Klammernprüfung
:set autoindent      automatische Einrückung
:set noautoindent    keine automatische Einrückung
:set wrapmargin=n    ab n Zeichen vor Zeilenende wird automatisch an einer Wortgrenze getrennt und eine neue Zeile begonnen
```

## ex 
Commands outside vi 

### Example 
``` 
$ vi HelloWorld.txt
Hallo Welt
:wq

$ vi HelloWorld.ex
" Change from German to Englisch words
%s/Hallo/Hello/g
%s/Welt/World/g
:wq

$ ex -s HelloWorld.txt < HelloWorld.ex
$ cat HelloWorld.txt
Hello World
``` 

## Links 
https://itler.net/linux-editor-vi-befehle-in-der-bersicht/

https://www-user.tu-chemnitz.de/~hot/VIM/vi.html

   
