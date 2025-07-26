---
title: Perché Blue ha una Beta Aperta
category: "Engineering"
description: Scoprite perché il nostro sistema di gestione progetti ha una beta aperta continua.
date: 2024-08-03
---

Molte startup B2B SaaS lanciano in Beta, e per buone ragioni. Fa parte del tradizionale motto della Silicon Valley *"muoviti velocemente e rompi le cose"*.

Mettere un'etichetta "beta" su un prodotto riduce le aspettative.

Qualcosa non funziona? Pazienza, è solo una beta.

Il sistema è lento? Pazienza, è solo una beta.

[La documentazione](https://blue.cc/docs) è inesistente? Pazienza... avete capito il concetto.

E questo è *effettivamente* una cosa positiva. Reid Hoffman, il fondatore di LinkedIN, disse notoriamente:

> Se non siete imbarazzati dalla prima versione del vostro prodotto, avete lanciato troppo tardi.

E l'etichetta beta è positiva anche per i clienti. Li aiuta ad auto-selezionarsi.

I clienti che provano prodotti beta sono quelli nelle prime fasi del Ciclo di Adozione Tecnologica, conosciuto anche come Curva di Adozione del Prodotto.

Il Ciclo di Adozione Tecnologica è tipicamente diviso in cinque segmenti principali:

1. Innovatori
2. Adottatori Precoci
3. Maggioranza Precoce
4. Maggioranza Tardiva
5. Ritardatari

![](/insights/technology-adoption-lifecycle-graph.png)

Tuttavia, alla fine il prodotto deve maturare e i clienti si aspettano un prodotto stabile e funzionante. Non vogliono accedere a un ambiente "beta" dove le cose si rompono.

O forse sì?

*Questa* è la domanda che ci siamo posti.

Crediamo di esserci posti questa domanda a causa della natura di come Blue è stato inizialmente costruito. [Blue è nato come un'emanazione di un'agenzia di design molto impegnata](/insights/agency-success-playbook), e quindi abbiamo lavorato *all'interno* dell'ufficio di un'azienda che utilizzava attivamente Blue per gestire tutti i suoi progetti.

Ciò significa che per anni siamo stati in grado di osservare come *esseri umani reali* — seduti proprio accanto a noi! — utilizzavano Blue nella loro vita quotidiana.

E poiché hanno utilizzato Blue fin dai primi giorni, questo team ha sempre utilizzato Blue Beta!

Ed è stato quindi naturale per noi permettere anche a tutti i nostri altri clienti di utilizzarlo.

**Ed è per questo che non abbiamo un team di test dedicato.**

Esatto.

Nessuno in Blue ha la *sola* responsabilità di garantire che la nostra piattaforma funzioni bene e sia stabile.

Questo per diverse ragioni.

La prima è una base di costi inferiore.

Non avere un team di test a tempo pieno riduce significativamente i nostri costi, e siamo in grado di trasferire questi risparmi ai nostri clienti con i prezzi più bassi del settore.

Per mettere questo in prospettiva, offriamo set di funzionalità di livello enterprise che la nostra concorrenza fa pagare $30-$55/utente/mese per soli $7/mese.

Questo non accade per caso, è *progettato*.

Tuttavia, non è una buona strategia vendere un prodotto più economico se non funziona.

Quindi la *vera domanda è*, come riusciamo a creare una piattaforma stabile che migliaia di clienti possono utilizzare senza un team di test dedicato?

Naturalmente, il nostro approccio di avere una Beta aperta è cruciale per questo, ma prima di approfondire questo aspetto, vogliamo toccare la responsabilità degli sviluppatori.

Abbiamo preso la decisione iniziale in Blue di non dividere mai le responsabilità per le tecnologie front-end e back-end. Avremmo assunto o formato solo sviluppatori full stack.

Il motivo per cui abbiamo preso questa decisione era di garantire che uno sviluppatore possedesse completamente la funzionalità su cui stava lavorando. Quindi non ci sarebbe stata quella mentalità del *"lanciare il problema oltre la recinzione del giardino"* che a volte si presenta quando ci sono responsabilità congiunte per le funzionalità.

E questo si estende al test della funzionalità, alla comprensione dei casi d'uso e delle richieste dei clienti, e alla lettura e commento delle specifiche.

In altre parole, ogni sviluppatore costruisce una comprensione profonda e intuitiva della funzionalità che sta costruendo.

Ok, parliamo ora della nostra beta aperta.

Quando diciamo che è "aperta" — lo intendiamo davvero. Qualsiasi cliente può provarla semplicemente aggiungendo "beta" davanti all'url della nostra applicazione web.

Quindi "app.blue.cc" diventa "beta.app.blue.cc"

Quando lo fanno, sono in grado di vedere i loro dati usuali, poiché sia gli ambienti Beta che Production condividono lo stesso database, ma potranno anche vedere nuove funzionalità.

I clienti possono lavorare facilmente anche se hanno alcuni membri del team su Production e alcuni curiosi su Beta.

Tipicamente abbiamo qualche centinaio di clienti che utilizzano Beta in qualsiasi momento, e pubblichiamo anteprime delle funzionalità sui nostri forum della community in modo che possano controllare le novità e provarle.

E questo è il punto: abbiamo *diverse centinaia* di tester!

Tutti questi clienti proveranno le funzionalità nei loro flussi di lavoro, e saranno piuttosto espliciti se qualcosa non va bene, perché stanno *già* implementando la funzionalità all'interno della loro azienda!

Il feedback più comune sono piccoli ma utilissimi cambiamenti che affrontano casi limite che non avevamo considerato.

Lasciamo le nuove funzionalità su Beta tra le 2-4 settimane. Ogni volta che sentiamo che sono stabili, le rilasciamo in produzione.

Abbiamo anche la possibilità di bypassare Beta se necessario, utilizzando un flag fast-track. Questo viene tipicamente fatto per correzioni di bug che non vogliamo trattenere per 2-4 settimane prima di spedire in produzione.

Il risultato?

Spingere in produzione si sente... beh, noioso! Come niente — semplicemente non è un grosso problema per noi.

E significa che questo uniforma il nostro programma di rilascio, che è ciò che ci ha permesso di [rilasciare funzionalità mensilmente come un orologio negli ultimi sei anni.](/changelog).

Tuttavia, come ogni scelta, ci sono alcuni compromessi.

Il supporto clienti è leggermente più complesso, poiché dobbiamo supportare i clienti su due versioni della nostra piattaforma. A volte questo può causare confusione ai clienti che hanno membri del team che utilizzano due versioni diverse.

Un altro punto dolente è che questo approccio può talvolta rallentare il programma di rilascio complessivo in produzione. Questo è particolarmente vero per funzionalità più grandi che possono rimanere "bloccate" in Beta se c'è un'altra funzionalità correlata che sta avendo problemi e necessita di ulteriore lavoro.

Ma nel complesso, pensiamo che questi compromessi valgano i benefici di una base di costi inferiore e di un maggiore coinvolgimento dei clienti.

Siamo una delle poche società di software ad abbracciare questo approccio, ma ora è una parte fondamentale del nostro approccio allo sviluppo del prodotto.