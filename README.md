# GoLang_Blog_Programmieren_II

Anforderungen
1. Nicht funktional
	1.1) Der Code soll im Paket de/vorlesung/projekt/[Gruppen-ID] liegen, damit alle Losungen parallel im Source-Tree gehalten werden k¨onnen.
	1.2) Es durfen keine Pakete Dritter verwendet werden! Einzige Ausnahme sind Pakete zur Vereinfachung der Tests und der Fehlerbehandlung. Empfohlen sei hier github.com/stretchr/testify/assert und github.com/pkg/errors.
	1.3) Alle Source-Dateien und die PDF-Datei mussen die Matrikelnummern aller Gruppenmitglieder enthalten.
2. Allgemein
	2.1) Die Anwendung muss nur ein Blog verwalten konnen. Mehrere Blogs auf demselben Server sind nicht erforderlich.
	2.2) Die Anwendung soll unter Windows und Linux lauﬀ¨ahig sein.
	2.3) Es soll sowohl IE 11 als auch Firefox, Chrome und Edge in der jeweils aktuellen Version unterstutzt werden. Diese Anforderung ist am einfachsten zu erfullen, indem Sie auf komplexe JavaScript/CSS ”spielereien“ verzichten und ein 90er Jahre Look&Feel in kauf nehmen. ;-)
3. Sicherheit
	3.1) Die Web-Seite soll nur per HTTPS erreichbar sein.
	3.2) Der Zugang fur die Autoren soll durch Benutzernamen und Passwort geschutzt werden.
	3.3) Die Passw¨orter durfen nicht im Klartext gespeichert werden.
	3.4) Es soll ”salting“ eingesetzt werden.
	3.5) Alle Zugangsdaten sind in einer gemeinsamen Datei zu speichern.
4. Ein Blog-Beitrag
	4.1) Ein Beitrag besteht aus dem eigentlichen Text,
	4.2) dem Datum der Erstellung,
	4.3) dem Namen des Autoren und
	4.4) den Kommentaren der Leser.
5. Ein Kommentar
	5.1) Ein Kommentar besteht aus dem eigentlichen Text,
	5.2) dem Datum der Erstellung und
	5.3) dem Nickname des Lesers.
6. WEB-Oberﬂache, Autoren
	6.1) Die Anmeldung soll uber eine Web-Seite erfolgen.
		• Zur weiteren Identiﬁkation des Nutzers soll ein Session-ID Cookie verwendet werden.
		• Der Session-Timeout soll 15 min betragen, und uber Flags einstellbar sein.
	6.2) Ein Autor soll sein Passwort ¨andern k¨onnen.
	6.3) Ein Autor soll einen Beitrag erstellen k¨onnen.
	6.4) Ein Autor soll nur seine eigenen Beitr¨age bearbeiteten und
	6.5) l¨oschen k¨onnen.
7. WEB-Oberﬂ¨ache, Leser
	7.1) Zun¨achst soll der neueste Beitrag zu sehen sein.
	7.2) Es soll m¨oglich sein, zu ¨alteren Beitr¨agen zu navigieren.
	7.3) Das Erstellungsdatum und der Autor des Beitrags soll ersichtlich sein.
	7.4) Kommentare sollen ebenfalls angezeigt werden.
	7.5) Anzeige nach Erstellungsdatum sortiert, neueste Kommentare oben
	7.6) Das Erstellungsdatum und der ”Nickname“ des Autors jedes Kommentars soll ersichtlich sein.
	7.7) Ein Leser soll einen neuen Kommentar abgeben k¨onnen.
	7.8) Eine Authentiﬁzierung ist nicht erforderlich.
	7.9) Es soll ein ”Nickname“ eingegeben werden k¨onnen.
8. Storage
	8.1) Jeder Beitrag soll incl. der Kommentare in einer Datei im Dateisystem gespeichert werden.
9. Konﬁguration
	9.1) Die Konﬁguration soll komplett uber Startparameter erfolgen. (Siehe Packageflag)
	9.2) Der Port muss sich uber ein Flag festlegen lassen.
	9.3) ”Hart kodierte“ absolute Pfade sind nicht erlaubt.
10. Betrieb
	10.1) Wird die Anwendung ohne Argumente gestartet, soll ein sinnvoller ”default“ gew¨ahlt werden.
	10.2) Nicht vorhandene aber ben¨otigte Order sollen ggﬂs. angelegt werden.
	10.3) Die Anwendung soll zwar HTTPS und die entsprechenden erforderlichen Zertiﬁkate unterstutzen, es kann jedoch davon ausgegangen werden, dass geeignete Zertiﬁkate gestellt werden. Fur Ihre Tests k¨onnen Sie ”self signed“ Zertiﬁkate verwenden. Es ist nicht erforderlich zur Laufzeit Zertiﬁkate zu erstellen o.¨a..