# Källskärsudden Community Web App

Detta är en community-webbapplikation för invånare som bor längs vägen "Källskärsudden". Applikationen är skriven med Ruby on Rails och Angular.js och syftar till att underlätta kommunikationen mellan grannar samt rapportering av hål i vägen längs Källskärsudden. Endast personer med fastighet längs vägen kan använda webbplatsen.

## Funktioner

- **Kommunikation**: Håll dig uppdaterad genom att kommunicera med grannar längs Källskärsudden. Dela information, ställ frågor och diskutera relevanta ämnen.

- **Rapportera hål i vägen**: Få möjlighet att rapportera hål i vägen genom att placera markörer på en karta. Hålen kan sedan verifieras av medlemmar när de åtgärdas.

## Installation

Följ dessa steg för att installera och köra applikationen lokalt:

1. Klona detta Git-repositorium till din lokala maskin:
  ```shell
  git clone https://github.com/<ditt användarnamn>/<ditt repo>.git
  ```
2. Gå till projektets rotmapp:
  ```shell
  cd road-community-ruby-angular
  ```
3. Installera alla beroenden med hjälp av bundler och npm:
  ```shell
  bundle install
  npm install
  ```
4. Konfigurera databasen genom att skapa databasen och köra migrations:
  ```shell
  rails db:create
  rails db:migrate
  ```
5. Starta servern:
  ```shell
  rails server
  ```
  
## Bidrag
Om du vill bidra till detta projekt, är du välkommen att skicka in en pull-begäran. 
Vi uppskattar alla bidrag och förbättringsförslag!

## Licens
Detta projekt är licensierat under MIT-licensen
