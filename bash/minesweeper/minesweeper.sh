#!/bin/bash

declare -a mine

function checkMargins() {
  if [ "${1}" -ge 0 -a "${2}" -ge 0 \
       -a "${1}" -lt "${w}" -a "${2}" -lt "${h}" ]; then
    return 0
  fi
  return 1
}

function mineCount() {
  local x y e=""
  local IFS=$'\n'
  for ((x=${1}-1; ${x}<=${1}+1; x++)) do
    for ((y=${2}-1; ${y}<=${2}+1; y++)) do
      e+=$(getField ${x} ${y})$'\n'
    done
  done
  grep -c "[*M]" <<< ${e}
}

function fieldCount() {
  local IFS=$'\n'
  grep -c "${1}" <<< ${mine[*]}
}

function getField() {
  if checkMargins "${1}" "${2}"; then
    echo "${mine[${2}*${w}+${1}]}"
  else
    echo "-"
  fi
}

function showField() {
  local x y i
  if checkMargins "${1}" "${2}"; then
    if [ "${mine[${2}*${w}+${1}]}" == '?' ]; then
      i=$(mineCount "${1}" "${2}")
      mine[${2}*${w}+${1}]="${i}"
    fi
  fi
}

function getMatrix() {
  local IFS=''
  echo "${mine[*]}"
}

function showMatrix() {
  local y
  for ((y=0; ${y}<${h}; y++)) do
    echo "${1:${y}*${w}:${w}}"
  done
}

function readCommands() {
  local x y line
  read -p '> ' line
  if [[ ${line} =~ ^[0-9]+\ [0-9]+$ ]]; then
    read x y <<< ${line}
    if checkMargins "$x" "$y"; then
      if [ "${mine[${y}*${w}+${x}]}" == '*' ]; then
        mine[${y}*${w}+${x}]='X'
        clear
        echo "Boom!"
        showMatrix "$(getMatrix)"
        exit 0
      else
        showField ${x} ${y}
      fi           
    else
      echo "Fehlerhafte Koordinate"
      read
    fi
  else
    echo "Befehl nicht erkannt"
    read
  fi
}

function showUsage() {
  echo "Aufruf: ${0} [-h] [-s <Breite>x<Höhe>] [-n <Minen>]"
}

while getopts ":hn:s:" opt; do
  case ${opt} in
    h)
      showUsage ${0}
      exit 0
      ;;
    n)
      if [[ ${OPTARG} =~ ^[0-9]+$ ]]; then
        n=${OPTARG}
      else
        showUsage ${0}
        exit 1
      fi
      ;;
    s)
      if [[ ${OPTARG} =~ ^[0-9]+x[0-9]+$ ]]; then
        w=${OPTARG%x*}
        h=${OPTARG#*x}
      else
        showUsage ${0}
        exit 1
      fi
      ;;
    \?)
      echo "Unbekannte Option -${OPTARG}"
      showUsage ${0}
      exit 1
  esac
done

: ${w:=8}
: ${h:=8}
: ${n:=10}

if [ $((${w}*${h})) -le ${n} ]; then
  echo "Kann nicht ${n} Minen auf ${w}x${h} Feldern verstecken."
  exit 1
fi

for ((i=0; ${i}<${w}*${h}; i++)) do
  mine[${i}]="?"
done

for i in $(shuf -i 0-$((${w}*${h}-1)) -n ${n}); do
  mine[$i]="*"
done

while true; do
  clear
  if [ $(fieldCount '?') == 0 ]; then
    echo "Geschafft!"
    showMatrix "$(getMatrix)"
    break
  else
    echo "Noch $(fieldCount '?') Feld(er)"
  fi
  m=$(getMatrix)
  m=${m//\*/?}
  m=${m//0/ }
  showMatrix "${m}" 
  readCommands
done

