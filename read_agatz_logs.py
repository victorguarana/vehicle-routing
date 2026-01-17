import re
import csv

# Função para processar o arquivo de log e gerar o CSV
def process_log_to_csv(log_file, output_csv):
    alpha = 'Alpha1'  # Valor fixo para a coluna Alpha

    # Regex para capturar as informações relevantes
    pattern = re.compile(
        r"AGATZ - Uniform - (N\d+) - (\d+) (\w+DecoderWithVehicleByPercentage) \(TimeSpent - Total Distance\): ([\d.]+)\n"
        r".*?TimeSpent - Time Spent\): ([\d.]+)"
    )

    # Lista para armazenar os dados extraídos
    data = []

    # Abrir o arquivo de log e processar linha por linha
    with open(log_file, 'r') as file:
        content = file.read()
        matches = pattern.findall(content)
        for match in matches:
            # Substituir ponto por vírgula para os valores numéricos
            n, instance, decoder, total_distance, time_spent = match
            total_distance = total_distance.replace('.', ',')
            time_spent = time_spent.replace('.', ',')
            data.append([n, instance, decoder, total_distance, time_spent, alpha])

    # Escrever os dados extraídos em um arquivo CSV
    with open(output_csv, 'w', newline='', encoding='utf-8') as csvfile:
        csvwriter = csv.writer(csvfile, delimiter=';')
        # Escrever o cabeçalho opcionalmente
        csvwriter.writerow(['N', 'Instance', 'Decoder', 'Total Distance', 'Time Spent', 'Alpha'])
        # Escrever os dados
        csvwriter.writerows(data)

    print(f"Arquivo CSV gerado com sucesso: {output_csv}")

# Exemplo de uso
log_file = "resultados_agatz_uniform_alpha1_05.log"  # Substitua pelo caminho do arquivo de log
output_csv = "resultados_agatz_uniform_05.csv"  # Nome do arquivo CSV de saída
process_log_to_csv(log_file, output_csv)