from flask import Flask, request, jsonify
import os
import subprocess
import uuid
import shutil

app = Flask(__name__)

SIMULATIONS_DIR = '/tmp/simulations'

@app.route('/simulate', methods=['POST'])
def simulate():
    """General purpose SPICE simulation endpoint"""
    netlist_content = request.get_data(as_text=True)
    if not netlist_content:
        return jsonify({'error': 'No netlist provided'}), 400

    os.makedirs(SIMULATIONS_DIR, exist_ok=True)

    sim_id = str(uuid.uuid4())
    sim_dir = os.path.join(SIMULATIONS_DIR, sim_id)
    os.makedirs(sim_dir)
    netlist_path = os.path.join(sim_dir, 'circuit.sp')

    try:
        # Write netlist file
        with open(netlist_path, 'w') as f:
            f.write(netlist_content)

        # Run simulation
        cmd = ["ngspice", "-b", "circuit.sp"]
        process = subprocess.run(cmd, cwd=sim_dir, capture_output=True, text=True)

        # Check for errors
        if process.returncode != 0:
            return jsonify({
                'status': 'failure',
                'ngspice_output': process.stdout,
                'ngspice_error': process.stderr,
                'return_code': process.returncode
            }), 500

        # Collect all output files
        output_files = {}
        for filename in os.listdir(sim_dir):
            if filename != 'circuit.sp':  # Skip the input netlist
                with open(os.path.join(sim_dir, filename), 'r') as f:
                    output_files[filename] = f.read()

        return jsonify({
            'status': 'success',
            'ngspice_output': process.stdout,
            'output_files': output_files
        })

    except Exception as e:
        return jsonify({
            'error': str(e),
            'type': type(e).__name__
        }), 500

    finally:
        # Clean up simulation directory
        shutil.rmtree(sim_dir, ignore_errors=True)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=int(os.environ.get('PORT', 5000)))
