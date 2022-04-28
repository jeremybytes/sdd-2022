..\go-demo\tour-of-go.exe > output.txt

if ($LASTEXITCODE -eq 0) { 
    Write-Host "Completed Successfully" 
} else { 
    Write-Host "ERROR!" 
}