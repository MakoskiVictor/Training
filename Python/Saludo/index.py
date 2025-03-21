from rich.console import Console

# Inicializo la consola de rich
console = Console()
# Pedir nombre
nombre = input('Hola, usuario, ¿cómo te llamas? \nNombre: ')

# Pedir edad
edad = input(f'¡Hola, {nombre}! ¿Cuántos años tienes? \nEdad: ')

# Manejador de errores por si no ingresa un número
try:
    # Convertir edad en un número
    edad = int(edad)
    if edad >= 18:
        print(f'Vaya, {nombre}, veo que eres mayor de edad.\n')
    else:
        print(f'Vaya, {nombre}, veo que eres menor de edad.\n')
except ValueError:
    console.print('Debías ingresar un número válido :(')
    exit()


continuar = console.input(prompt='¿Te gustaría seguir aprendiendo? (Sí/No) \n>')

if continuar.lower() in ('sí', 'si', 's'):
    console.print('¡Perfecto, continuemos en la siguiente App!')
elif continuar.lower() in ('no', 'n'):
    console.print('Una lástima. Nos vemos :(')
else:
    console.print('Lo lamento, no entiendo tu respuesta. Por favor, solo responde con Sí o No')
