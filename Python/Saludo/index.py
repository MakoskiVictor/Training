import time
from rich.console import Console

class App:
    # Inicializo la consola de rich
    def __init__(self):
        self.console = Console()
        self.nombre = None
        self.edad = None
        self.continuar = None

    # Pedir nombre
    def pedir_nombre(self):
        self.nombre = self.console.input('[bold green]Hola, usuario, ¿cómo te llamas? \n[bold yellow]Nombre: ')

    # Pedir edad
    def pedir_edad(self):
        self.edad = self.console.input(f'[bold green]¡Hola, {self.nombre}! ¿Cuántos años tienes? \n[bold yellow]Edad: ')


    def mayor_edad(self):
        # Manejador de errores por si no ingresa un número
        try:
            # Convertir edad en un número
            self.edad = int(self.edad)
            if self.edad >= 18:
                self.console.print(f'[bold green]Vaya, {self.nombre}, veo que eres mayor de edad.\n')
            else:
                self.console.print(f'[bold green]Vaya, {self.nombre}, veo que eres menor de edad.\n')
        except ValueError:
            self.error_num()

    # Error en la edad
    def error_num(self):
        self.console.print('[bold red]Debías ingresar un número válido :(')
        self.pedir_edad()
        self.mayor_edad()

    # Continuar aprendiendo
    def continuar_aprendiendo(self):
        continuar = self.console.input(prompt='[bold green]¿Te gustaría seguir aprendiendo? (Sí/No) \n[bold yellow]>')

        if continuar.lower() in ('sí', 'si', 's'):
            self.console.print('[bold green]¡Perfecto, continuemos en la siguiente App!')
            time.sleep(5)
            exit()
        elif continuar.lower() in ('no', 'n'):
            self.console.print('[bold red]Una lástima. Nos vemos :(')
            time.sleep(5)
            exit()
        else:
            self.console.print('[bold red]Lo lamento, no entiendo tu respuesta. Por favor, solo responde con Sí o No')
            self.continuar_aprendiendo()

    # Manejador de la App
    def run(self):
        self.pedir_nombre()
        self.pedir_edad()
        self.mayor_edad()
        self.continuar_aprendiendo()

app = App()
app.run()
