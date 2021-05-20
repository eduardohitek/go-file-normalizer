package main

import "testing"

func Test_normalizeFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Remover acentos minúsculos", args: args{fileName: "joãofuraçãoéíó.txt"}, want: "joaofuracaoeio.txt"},
		{name: "Remover acentos maiúsculos", args: args{fileName: "JOÃOFURAÇÃOÉÍÓ.txt"}, want: "JOAOFURACAOEIO.txt"},
		{name: "Trocar espaços por underline", args: args{fileName: "nome do arquivo.txt"}, want: "nome_do_arquivo.txt"},
		{name: "Todas as regras anteriores", args: args{fileName: "JOÃO furação áéíóú ÁÉÍÓÚ.txt"}, want: "JOAO_furacao_aeiou_AEIOU.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeFileName(tt.args.fileName); got != tt.want {
				t.Errorf("normalizeFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
