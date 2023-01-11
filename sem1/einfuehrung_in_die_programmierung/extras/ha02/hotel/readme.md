# Abgabe zur Hausaufgabe 2
Erarbeitet von Paul Friedrich Vierkorn

### Notizen zu Aufgabe 1

Das Programm arbeitet intern mit sog. Points of Interest, welche im Code und in dem folgenden Text als POI abgekürzt werden. Ein POI besteht dabei aus einem Item, einer nextID und einer nextDistance. Die nextID wird vom aktuellen Programm nicht verwendet, kann aber in der Zukunft verwendet werden, um nichtlineare Abläufe zwischen den Hotels zu ermöglichen.
Durch Flags, die im POI-Struct hinterlgt werden könnten, könnten zum Beispiel auch Landschaftsmarken, Museen und andere Standorte ins Dataset hinzugefügt werden, um die Routenführung für Reisende interessanter zu gestalten.

Items speichern ihren eienenen Namen und eine Description.
Diese Description wird als Teil der Waypoints nach der Wegeberechnung ausgegeben und bietet eine kurze Beschreibung zu POIs.

Das Einlesen von Daten aus der CSV-Datei wird hierbei dynamisch gestaltet. Ist eine Description im zweiten Feld vorhanden, wird die Description zuerst geladen.

### Notizen zu Aufgabe 2

Um die Nutzung des Tools angenehmer zu gestalten, wird nicht nur die Distanz zwischen den beiden POIs ausgegeben, sondern auch die Reihenfolge, welche POIs angefaahren werden sollten.  
Somit wird verhindert, dass die Radfahrer in eine falsche Richtung losfahren.

Hierbei wird ebenfalls die Item.Description ausgegeben, die einzelne POIs auf der Strecke beschreibt.

### Notizen zu Aufgabe 3

Besonders bei Rasterbildformaten ist das Zeichnen von Bildern über Bitmaps einfach umsetzbar. 
Da sich Bitmaps gut eignen um Pixelbasierte Formen zu speichern, wurde in der Datei bitfont.map in der imaging Library eine Schriftart (Font) definiert, welche durch ihre kleine Schriftgröße super zur Bildausgabe passt.

Die bitfont.map-Datei, welche eine Plaintextdatei ist, wird vom Pythonskript GoFontMaker.py automatisiert in eine golang Map umgewandelt und in die Datei bitfont.go geschrieben.

Anschließend kann die Font durch die Imaging-Library verwendet werden, um einzelne Bitmaps auf das Endbild zu bringen.  
Da jeder Rune eine Bitmap (in diesem Falle genaugenommen eine Bool-Map) zugeordnet ist, kann nun einfach Text auf das Bild gebracht werden.

Das Programm sucht hierzu erst nach den gewünschten POIs, erstellt das Bild, erzeugt darauf Icons, und zeichnet mehrfach Text auf das Bild.

Ausgegeben werden hier der Name beider Hotels, der Mindestabstand, sowie die Anzahl der POIs, die auf der Route liegen.