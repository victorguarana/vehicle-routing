import os

# Caminho da pasta com os arquivos
input_folder = "/home/victor/projetos/vehicle-routing/TSP-D-Instances/doublecenter"
output_folder = "/home/victor/projetos/vehicle-routing/cmd/execute/instances/agatz/doublecenter"

# Função para processar cada arquivo
def process_file(input_path, output_path):
    with open(input_path, "r") as infile:
        lines = infile.readlines()
    
    # Extraindo informações do depósito
    depot_line = lines[7].strip().split()
    depot_x, depot_y = depot_line[0], depot_line[1]
    depot = f"dep;{depot_x};{depot_y};0"
    
    # Extraindo informações das locations
    locations = []
    for line in lines[8:]:
        parts = line.strip().split()
        if len(parts) == 3:
            x, y, name = parts
            locations.append(f"\n{name};{x};{y};1")
    
    # Escrevendo no novo arquivo
    with open(output_path, "w") as outfile:
        outfile.write(depot)
        outfile.writelines(locations)

def should_process_file(filename):
    return not filename.startswith("doublecenter-alpha") and (filename.endswith("n10.txt") or filename.endswith("n20.txt") or filename.endswith("n50.txt") or filename.endswith("n75.txt") or filename.endswith("n100.txt") or filename.endswith("n175.txt") or filename.endswith("n250.txt"))


# Processando todos os arquivos na pasta
for filename in os.listdir(input_folder):
    if should_process_file(filename):
        input_path = os.path.join(input_folder, filename)
        output_filename = filename.replace(".txt", "")
        output_path = os.path.join(output_folder, output_filename)
        process_file(input_path, output_path)

print("Arquivos convertidos com sucesso!")