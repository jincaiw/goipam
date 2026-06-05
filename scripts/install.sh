#!/bin/bash
set -e
# Install GoIPAM
INSTALL_DIR="/opt/goipam"
sudo useradd --system --home /var/lib/goipam --shell /usr/sbin/nologin goipam 2>/dev/null || true
sudo install -d -o goipam -g goipam -m 0750 "$INSTALL_DIR" /var/lib/goipam /etc/goipam
sudo install -m 0755 bin/goipam "$INSTALL_DIR/goipam"
sudo install -m 0640 configs/config.example.yaml "$INSTALL_DIR/config.yaml"
if [ ! -f /etc/goipam/goipam.env ]; then
  sudo install -m 0640 /dev/null /etc/goipam/goipam.env
  echo "GOIPAM_ENV=production" | sudo tee /etc/goipam/goipam.env >/dev/null
  echo "GOIPAM_ADMIN_PASSWORD=change-this-before-start" | sudo tee -a /etc/goipam/goipam.env >/dev/null
fi
sudo chown root:goipam /etc/goipam/goipam.env
sudo cp scripts/goipam.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable goipam
echo "GoIPAM installed. Edit $INSTALL_DIR/config.yaml and /etc/goipam/goipam.env, then run: sudo systemctl start goipam"
