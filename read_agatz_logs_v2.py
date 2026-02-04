import re
import csv
import sys

# Função para processar o arquivo de log e gerar o CSV
def process_log_to_csv(log_file, output_csv, alpha):
    # Regex para capturar as informações relevantes
    pattern = re.compile(
        r"AGATZ - (Uniform|SingleCenter|DoubleCenter) - (N\w+) - (\d+) (\w+DecoderWithVehicleByPercentage) \(TimeSpent - Time Spent\): ([\d.]+)"
    )

    # Lista para armazenar os dados extraídos
    data = []

    # Abrir o arquivo de log e processar linha por linha
    with open(log_file, 'r') as file:
        content = file.read()
        matches = pattern.findall(content)
        for match in matches:
            # Substituir ponto por vírgula para os valores numéricos
            _, n, instance, decoder, time_spent = match
            time_spent = time_spent.replace('.', ',')
            data.append([n, instance, decoder, time_spent, alpha])

    # Escrever os dados extraídos em um arquivo CSV
    with open(output_csv, 'w', newline='', encoding='utf-8') as csvfile:
        csvwriter = csv.writer(csvfile, delimiter=';')
        # Escrever o cabeçalho opcionalmente
        csvwriter.writerow(['N', 'Instance', 'Decoder', 'Time Spent', 'Alpha'])
        # Escrever os dados
        csvwriter.writerows(data)

    print(f"Arquivo CSV gerado com sucesso: {output_csv}")

# Exemplo de uso

# Recebendo arquivo por argumento de linha de comando
alpha = sys.argv[2]  # Substitua pelo caminho do arquivo de log
log_file = sys.argv[1] + ".log"  # Substitua pelo caminho do arquivo de log
output_csv = sys.argv[1] + ".csv"  # Nome do arquivo CSV de saída
process_log_to_csv(log_file, output_csv, alpha)