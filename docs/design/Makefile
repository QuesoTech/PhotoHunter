TARGET=design

all:
	pdflatex $(TARGET)
	pdflatex $(TARGET)
	pdflatex $(TARGET)

spell:
	aspell -ttex -c $(TARGET).tex

clean:
	rm -f $(TARGET).pdf *.blg *.aux *.bbl *.brf *.log *.nav *.snm *.toc *.out
