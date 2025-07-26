---
title: Ricerca in tempo reale
description: Blue svela un nuovo motore di ricerca super veloce che restituisce risultati in tutti i tuoi progetti in millisecondi, permettendoti di cambiare contesto in un batter d'occhio.
category: "Product Updates"
date: 2024-03-01
---


Siamo entusiasti di annunciare il lancio del nostro nuovo motore di ricerca, progettato per rivoluzionare il modo in cui trovi informazioni all'interno di Blue. Una funzionalità di ricerca efficiente è fondamentale per una gestione dei progetti senza soluzione di continuità, e il nostro ultimo aggiornamento garantisce che tu possa accedere ai tuoi dati più velocemente che mai.

Il nostro nuovo motore di ricerca ti consente di cercare tutti i commenti, i file, i registri, i campi personalizzati, le descrizioni e le checklist. Che tu abbia bisogno di trovare un commento specifico fatto su un progetto, localizzare rapidamente un file o cercare un particolare record o campo, il nostro motore di ricerca fornisce risultati fulminei.

Man mano che gli strumenti si avvicinano a una reattività di 50-100 ms, tendono a scomparire e a fondersi con lo sfondo, offrendo un'esperienza utente senza soluzione di continuità e quasi invisibile. Per contestualizzare, un battito di ciglia umano richiede circa 60-120 ms, quindi 50 ms è in realtà più veloce di un battito di ciglia! Questo livello di reattività ti consente di interagire con Blue senza nemmeno renderti conto che è lì, liberandoti per concentrarti sul lavoro effettivo. Sfruttando questo livello di prestazioni, il nostro nuovo motore di ricerca garantisce che tu possa accedere rapidamente alle informazioni di cui hai bisogno, senza che ciò interferisca mai con il tuo flusso di lavoro.

Per raggiungere il nostro obiettivo di una ricerca fulminea, abbiamo sfruttato le ultime tecnologie open-source. Il nostro motore di ricerca è costruito su MeiliSearch, un popolare servizio di ricerca open-source che utilizza l'elaborazione del linguaggio naturale e la ricerca vettoriale per trovare rapidamente risultati pertinenti. Inoltre, abbiamo implementato uno storage in memoria, che ci consente di memorizzare i dati frequentemente accessibili nella RAM, riducendo il tempo necessario per restituire i risultati della ricerca. Questa combinazione di MeiliSearch e storage in memoria consente al nostro motore di ricerca di fornire risultati in millisecondi, rendendo possibile per te trovare rapidamente ciò di cui hai bisogno senza dover mai pensare alla tecnologia sottostante.

La nuova barra di ricerca è comodamente posizionata sulla barra di navigazione, permettendoti di iniziare a cercare subito. Per un'esperienza di ricerca più dettagliata, basta premere il tasto Tab mentre cerchi per accedere alla pagina di ricerca completa. Inoltre, puoi attivare rapidamente la funzione di ricerca da qualsiasi luogo utilizzando la scorciatoia CMD/Ctrl+K, rendendo ancora più facile trovare ciò di cui hai bisogno.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>

## Sviluppi Futuri

Questo è solo l'inizio. Ora che abbiamo un'infrastruttura di ricerca di nuova generazione, possiamo fare alcune cose davvero interessanti in futuro.

Il prossimo passo sarà la ricerca semantica, che rappresenta un miglioramento significativo rispetto alla tipica ricerca per parole chiave. Permettici di spiegare.

Questa funzionalità permetterà al motore di ricerca di comprendere il contesto delle tue query. Ad esempio, cercare "mare" restituirà documenti pertinenti anche se non viene utilizzata l'esatta frase. Potresti pensare "ma ho digitato 'oceano' invece!" - e hai ragione. Il motore di ricerca comprenderà anche la somiglianza tra "mare" e "oceano", e restituirà documenti pertinenti anche se non viene utilizzata l'esatta frase. Questa funzionalità è particolarmente utile quando si cercano documenti contenenti termini tecnici, acronimi o semplicemente parole comuni che hanno più variazioni o errori di battitura.

Un'altra funzionalità in arrivo è la possibilità di cercare immagini in base al loro contenuto. Per raggiungere questo obiettivo, elaboreremo ogni immagine nel tuo progetto, creando un embedding per ciascuna. In termini generali, un embedding è un insieme matematico di coordinate che corrisponde al significato di un'immagine. Ciò significa che tutte le immagini possono essere cercate in base a ciò che contengono, indipendentemente dal loro nome file o dai metadati. Immagina di cercare "diagramma di flusso" e trovare tutte le immagini relative ai diagrammi di flusso, *indipendentemente dai loro nomi file.*