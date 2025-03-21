# Pedir nombre
nombre = input('Hola, usuario, ¿cómo te llamas? \nNombre: ')

# Pedir edad
edad = input(f'¡Hola, {nombre}! ¿Cuántos años tienes? \nEdad: ')

# Convertir edad en un número
edad = int(edad)

if edad >= 18:
    print(f'Vaya, {nombre}, veo que eres mayor de edad.')
else:
    print(f'Vaya, {nombre}, veo que eres menor de edad.')