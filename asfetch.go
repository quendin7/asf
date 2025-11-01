package main

import (
	"asf/config"
	"asf/desktop"
	"asf/dodatki"
	"asf/hardware"
	"asf/osinfo"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// KOLORY - zmienione na kolory ANSI, które są częścią standardowej palety terminala
var (
	ColorLabelLight = "\033[94m" // Jasny Cyan
	ColorLabelDark  = "\033[36m" // Ciemny Cyan
	ColorSepLight   = "\033[90m" // Jasny niebieski
	ColorSepDark    = "\033[97m" // Ciemny niebieski
	ColorValueLight = "\033[96m" // Jasny biały
	ColorValueDark  = "\033[34m" // Ciemny szary
	ColorReset      = "\033[0m"
)

func main() {
	allFlag := flag.Bool("a", false, "Włącz wszystkie opcje niezależnie od konfiguracji")
	logooffFlag := flag.Bool("l", false, "Wyłącz logo")
	onlylogoFlag := flag.Bool("L", false, "Wyświetl tylko logo")
	noneFlag := flag.Bool("b", false, "Włącz tryb przyjazny dla screen readerów (bez kolorów i logo)")
	ttsFlag := flag.Bool("tts", false, "Pokazuje parasol z That's the Spirit zamiast loga")
	defaultConfigFlag := flag.Bool("c", false, "Użyj domyślnego configu")

	flag.Parse()
	cfg := config.LoadConfig()
	if *allFlag {
		cfg = config.GetAllEnabledConfig()
	}

	if *defaultConfigFlag {
		cfg = config.GetDefaultConfig()
	}

	// obsługa trybu TTS (zastępuje logo własnym artem)
	useTTS := false
	var ttsLogoLines []string

	if *ttsFlag {
		cfg.EnableLogo = true
		useTTS = true
		ttsLogoLines = []string{
			"                       @+                       ",
			"                       @@                       ",
			"              :@@@@@@@@@@@@@@@@@@               ",
			"           @@@%   @@@      @@#   @@@#           ",
			"         @@=     @@          @@     @@@         ",
			"       @@       @@            @@      @@@       ",
			"     %@@        @             @@        @@      ",
			"    @@         @@              @@        %@=    ",
			"   @@          @@              @@         :@+   ",
			"  @@           @@              @@          *@   ",
			"  @.%@@@@@@@:  @@   @@@@@@@@   @@  @@@@@@@@ @@  ",
			" @@@#       @@@@@@@@   @@   @@ @@@@+       @@@@ ",
			"*@@           @@@@     @@    :@@@@           @@ ",
			"@@             @@      @@     %@@             @@",
			"@*             @@      @@      @.             @@",
			"        @              @@               @       ",
			"       @@              @@               @@      ",
			"        @              @@               @.      ",
			"                 @@    @@        @              ",
			"                 @@    @@        @@             ",
			"                       @@       %@*             ",
			"                       @@                       ",
			"                       @@                       ",
			"                       @@                       ",
		}
	}

	infoPairs := []struct {
		Label string
		Value string
	}{}
	if *noneFlag {
		cfg.EnableLogo = false
		cfg.EnableUserHost = true
		cfg.EnableOSInfo = true
		cfg.EnableKernel = true
		cfg.EnablePackages = true
		cfg.EnableDEWM = true
		cfg.EnableCPU = true
		cfg.EnableGPU = true
		cfg.EnableRAM = true
		cfg.EnableSwap = true
		cfg.EnableMusic = true
		cfg.EnableUptime = true
		cfg.EnableGTKTheme = false
		cfg.EnableIconTheme = false
		cfg.EnableFont = false
		cfg.EnableShell = false
		cfg.EnableBattery = true
		cfg.EnableDiskInfo = true
		cfg.EnableNetworkInfo = true
		cfg.EnableScreenResolution = false
		ColorLabelLight = ""
		ColorLabelDark = ""
		ColorSepLight = ""
		ColorSepDark = ""
		ColorValueLight = ""
		ColorValueDark = ""
		ColorReset = ""
	}
	if *logooffFlag {
		cfg.EnableLogo = false
	}
	if *onlylogoFlag {
		cfg.EnableLogo = true
		cfg.EnableUserHost = false
		cfg.EnableOSInfo = false
		cfg.EnableKernel = false
		cfg.EnablePackages = false
		cfg.EnableDEWM = false
		cfg.EnableCPU = false
		cfg.EnableGPU = false
		cfg.EnableRAM = false
		cfg.EnableSwap = false
		cfg.EnableMusic = false
		cfg.EnableUptime = false
		cfg.EnableGTKTheme = false
		cfg.EnableIconTheme = false
		cfg.EnableFont = false
		cfg.EnableShell = false
		cfg.EnableBattery = false
		cfg.EnableDiskInfo = false
		cfg.EnableNetworkInfo = false
		cfg.EnableScreenResolution = false
	}
	if useTTS {
		cfg.EnableLogo = true
	}
	if *onlylogoFlag && useTTS {
		cfg.EnableLogo = true
		cfg.EnableUserHost = false
		cfg.EnableOSInfo = false
		cfg.EnableKernel = false
		cfg.EnablePackages = false
		cfg.EnableDEWM = false
		cfg.EnableCPU = false
		cfg.EnableGPU = false
		cfg.EnableRAM = false
		cfg.EnableSwap = false
		cfg.EnableMusic = false
		cfg.EnableUptime = false
		cfg.EnableGTKTheme = false
		cfg.EnableIconTheme = false
		cfg.EnableFont = false
		cfg.EnableShell = false
		cfg.EnableBattery = false
		cfg.EnableDiskInfo = false
		cfg.EnableNetworkInfo = false
		cfg.EnableScreenResolution = false
		fmt.Printf("%s%swell, that's the spirit%s\n", ColorLabelLight, ColorValueLight, ColorReset)
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}

	if cfg.EnableUserHost {
		username, hostname := dodatki.GetUserAndHost()
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"User ", username + "@" + hostname})
	}

	if cfg.EnableOSInfo {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"OS ", osinfo.GetOSInfo()})
	}
	if cfg.EnableKernel {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"Kernel ", osinfo.GetKernel()})
	}
	if cfg.EnablePackages {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"Packages ", osinfo.GetPackageCount()})
	}

	if cfg.EnableDEWM {
		de, wm := desktop.GetDEWM()
		if de != "" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"DE ", de})
		}
		if wm != "" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"WM ", wm})
		}
	}

	if cfg.EnableGTKTheme {
		gtkTheme := desktop.GetGTKTheme()
		if gtkTheme != "unknown" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"GTK ", gtkTheme})
		}
	}
	if cfg.EnableIconTheme {
		iconTheme := desktop.GetIconTheme()
		if iconTheme != "unknown" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Icons ", iconTheme})
		}
	}
	if cfg.EnableFont {
		font := desktop.GetFont()
		if font != "unknown" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Font ", font})
		}
	}
	if cfg.EnableShell {
		shell := dodatki.GetShell()
		if shell != "unknown" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Shell ", shell})
		}
	}
	if cfg.EnableUptime {
		uptime := dodatki.GetUptime()
		if uptime != "unknown" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Uptime ", uptime})
		}
	}
	if cfg.EnableBattery {
		batteryInfo := hardware.GetBatteryInfo()
		if batteryInfo != "N/A" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Battery ", batteryInfo})
		}
	}
	if cfg.EnableDiskInfo {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"/  ", hardware.GetDiskInfo()})
	}
	if cfg.EnableNetworkInfo {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"Network ", hardware.GetNetworkInfo()})
	}
	if cfg.EnableScreenResolution {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"Resolution ", desktop.GetScreenResolution()})
	}
	if cfg.EnableCPU {
		infoPairs = append(infoPairs, struct {
			Label string
			Value string
		}{"CPU ", hardware.GetCPUInfo()})
	}
	if cfg.EnableGPU {
		gpuInfo := hardware.GetGPUInfo()
		if gpuInfo != "Unknown GPU" && gpuInfo != "N/A" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"GPU ", gpuInfo})
		}
	}

	if cfg.EnableRAM || cfg.EnableSwap {
		ramInfo, swapInfo := hardware.GetMemoryAndSwapInfo()
		if cfg.EnableRAM {
			if ramInfo != "unknown" {
				infoPairs = append(infoPairs, struct {
					Label string
					Value string
				}{"RAM ", ramInfo})
			}
		}
		if cfg.EnableSwap {
			if swapInfo != "unknown" && swapInfo != "No Swap" {
				infoPairs = append(infoPairs, struct {
					Label string
					Value string
				}{"Swap ", swapInfo})
			}
		}
	}

	if cfg.EnableMusic {
		musicInfo := dodatki.GetCurrentMusic()
		if musicInfo != "Not playing" {
			infoPairs = append(infoPairs, struct {
				Label string
				Value string
			}{"Spotify ", musicInfo})
		}
	}

	maxLabelLen := 0
	for _, pair := range infoPairs {
		if len(pair.Label) > maxLabelLen {
			maxLabelLen = len(pair.Label)
		}
	}

	var infoLines []string
	for i, pair := range infoPairs {
		var labelColor, sepColor, valueColor string
		if i%2 == 0 {
			labelColor = ColorLabelLight
			sepColor = ColorSepDark
			valueColor = ColorValueLight
		} else {
			labelColor = ColorLabelDark
			sepColor = ColorSepLight
			valueColor = ColorValueDark
		}
		alignedLabel := fmt.Sprintf("%s%-*s%s", labelColor, maxLabelLen, pair.Label, ColorReset)
		separator := fmt.Sprintf("%s│%s", sepColor, ColorReset)
		value := fmt.Sprintf("%s%s%s", valueColor, pair.Value, ColorReset)
		infoLines = append(infoLines, fmt.Sprintf("%s%s %s", alignedLabel, separator, value))
	}

	var logo []string
	if useTTS {
		logo = ttsLogoLines
	} else if cfg.EnableLogo {
		var err error
		logo, err = config.LoadLogoFromFile(cfg.LogoPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Błąd wczytywania logo z pliku: %v. Wyłączam logo.\n", err)
			logo = make([]string, len(infoLines))
			if len(infoLines) < 19 {
				logo = make([]string, 19)
			}
		}
	} else {
		logo = make([]string, len(infoLines))
		if len(infoLines) < 19 {
			logo = make([]string, 19)
		}
	}

	maxLogoWidth := 0
	if cfg.EnableLogo {
		for _, line := range logo {
			lineLen := 0
			for _, r := range line {
				if r >= 0x1F600 && r <= 0x1F64F {
					lineLen += 2
				} else if r >= 0x2000 && r <= 0x206F {
					lineLen += 1
				} else if r >= 0x2500 && r <= 0x257F {
					lineLen += 1
				} else if r >= 0x2800 && r <= 0x28FF {
					lineLen += 2
				} else if r >= 0x0000 && r <= 0x007F {
					lineLen += 1
				} else {
					lineLen += 1
				}
			}
			if lineLen > maxLogoWidth {
				maxLogoWidth = lineLen
			}
		}
	}
	if maxLogoWidth == 0 {
		maxLogoWidth = 20
	}

	totalHeight := len(logo)
	if len(infoLines) > totalHeight {
		totalHeight = len(infoLines)
	}

	emptyLinesTop := 1
	for i := 0; i < emptyLinesTop; i++ {
		fmt.Printf("%s\n", strings.Repeat(" ", maxLogoWidth))
	}

	for i := 0; i < totalHeight; i++ {
		logoLine := ""
		if i < len(logo) {
			logoLine = logo[i]
		}

		infoLine := ""
		if i < len(infoLines) {
			infoLine = infoLines[i]
		}

		calculatedWidth := 0
		if cfg.EnableLogo {
			for _, r := range logoLine {
				if r == ' ' {
					calculatedWidth += 1
				} else if r >= 0x2800 && r <= 0x28FF {
					calculatedWidth += 2
				} else if r >= 0x2580 && r <= 0x259F {
					calculatedWidth += 1
				} else {
					calculatedWidth += 1
				}
			}
		}

		spacing := 4
		if cfg.EnableLogo {
			fmt.Printf("%s%s%s%s%s%s\n", ColorValueLight, logoLine, ColorReset, strings.Repeat(" ", maxLogoWidth-calculatedWidth+spacing), infoLine, ColorReset)
		} else {
			fmt.Printf("%s%s\n", strings.Repeat(" ", maxLogoWidth+spacing), infoLine)
		}
	}

	fmt.Printf("%s\n", strings.Repeat(" ", maxLogoWidth))
}
