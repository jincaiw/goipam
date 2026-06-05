#!/bin/bash
set -e
BACKUP_DIR="/opt/goipam/backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
mkdir -p $BACKUP_DIR
# SQLite backup
cp /opt/goipam/data/goipam.db "$BACKUP_DIR/goipam_$TIMESTAMP.db"
# Keep last 30 backups
ls -t $BACKUP_DIR/goipam_*.db | tail -n +31 | xargs -r rm
echo "Backup created: $BACKUP_DIR/goipam_$TIMESTAMP.db"
