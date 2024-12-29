import numpy as np
import matplotlib.pyplot as plt

# Path to the file
file_path = "/Users/nolantremelling/PumpkinSpice/simulations/result.txt"

# Initialize lists to store data
frequencies = []
magnitudes = []

try:
    with open(file_path, "r") as file:
        lines = file.readlines()

    # Skip header lines (assume header is 6 lines)
    data_lines = lines[6:]

    # Process the lines
    for i in range(0, len(data_lines), 2):
        # Debugging: Print the lines being processed
        print(f"Processing line {i}: {data_lines[i].strip()}")

        try:
            # Extract frequency
            freq = float(data_lines[i].split("\t")[1].split(",")[0])  # Frequency
            # Extract real and imaginary parts of v(vout+)
            re, im = map(float, data_lines[i + 1].split("\t")[1].split(","))
            # Calculate magnitude
            magnitude = np.sqrt(re**2 + im**2)
            # Append to lists
            frequencies.append(freq)
            magnitudes.append(magnitude)
        except (IndexError, ValueError) as e:
            print(f"Error processing line {i}: {e}")
            continue  # Skip problematic lines

    # Convert magnitudes to dB
    magnitudes_db = 20 * np.log10(magnitudes)

    # Plot the data
    plt.figure(figsize=(10, 6))
    plt.semilogx(frequencies, magnitudes_db, label="Magnitude (dB)")
    plt.title("Low Pass Filter Analysis (10kHz)")
    plt.xlabel("Frequency (Hz)")
    plt.ylabel("Magnitude (dB)")
    plt.grid(True, which="both", linestyle="--", linewidth=0.5)
    plt.legend()
    plt.show()

except FileNotFoundError:
    print(f"File not found: {file_path}")
except Exception as e:
    print(f"Unexpected error: {e}")
