# Herencia Vs Composición

## Herencia 'ES UN'

Character <= Monster <= Orc
                    <= Dragon
          <= Hero

- Un Orc es un Monster que es un Character

- Hay jerarquías / niveles

- Polimorfismo a través de la herencia

- Test difíciles

## Composición 'TIENE / USA'

Bike <= Seat
    <= Pedals
    <= Chain

- Un Bike tiene un Seat y un Pedals y un Chain

- Hay relaciones y responsabilidades (La clase Bike le da la responsabilidad a Seat para que haga lo que debe hacer)

- Polimorfismo a través de interfaces

- Test sencillos
