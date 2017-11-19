
if expr "$1" : "^v[0-9]\.[0-9][0-9]*$" 1>/dev/null; then
  echo "Building $1";
else
  echo "Bad semver: \"$1\". Should be like \"v8.12\".";
  exit;
fi

git tag $1

# GOOS=windows GOARCH=386 go build -o ./voiceworkout-windows-$1.exe
# mv ./voiceworkout-windows-$1.exe releases
#
GOOS=linux GOARCH=386 go build -o ./voiceworkout-linux-$1
mv ./voiceworkout-linux-$1 releases

go build -o voiceworkout-mac-$1
mv ./voiceworkout-mac-$1 releases
