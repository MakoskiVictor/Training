import time
from rich.console import Console

# Inicializo la consola de rich
console = Console()
# Pedir nombre
nombre = console.input('[bold green]Hola, usuario, ¿cómo te llamas? \n[bold yellow]Nombre: ')

# Pedir edad
edad = console.input(f'[bold green]¡Hola, {nombre}! ¿Cuántos años tienes? \n[bold yellow]Edad: ')

# Manejador de errores por si no ingresa un número
try:
    # Convertir edad en un número
    edad = int(edad)
    if edad >= 18:
        console.print(f'[bold green]Vaya, {nombre}, veo que eres mayor de edad.\n')
    else:
        console.print(f'[bold green]Vaya, {nombre}, veo que eres menor de edad.\n')
except ValueError:
    console.print('[bold red]Debías ingresar un número válido :(')
    time.sleep(5)
    exit()


continuar = console.input(prompt='¿Te gustaría seguir aprendiendo? (Sí/No) \n[bold yellow]>')

if continuar.lower() in ('sí', 'si', 's'):
    console.print('[bold green]¡Perfecto, continuemos en la siguiente App!')
elif continuar.lower() in ('no', 'n'):
    console.print('[bold red]Una lástima. Nos vemos :(')
else:
    console.print('[bold red]Lo lamento, no entiendo tu respuesta. Por favor, solo responde con Sí o No')
time.sleep(5)
exit()