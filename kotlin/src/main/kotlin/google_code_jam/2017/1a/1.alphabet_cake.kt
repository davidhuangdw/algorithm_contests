package google_code_jam.`2017`.`1a`
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (T, ) = input_ints()
    (1..T).forEach {
        val (R, C) = input_ints()
        val A = (1..R).map{split_input()[0].toMutableList()}.toMutableList()
        for(row in A){
            var exist = false
            for((i, v) in row.withIndex()){
                if(v == '?'){
                    if(exist)
                        row[i] = row[i-1]
                }else{
                    if(!exist){
                        exist = true
                        for(j in 0 until i)
                            row[j] = row[i]
                    }
                }
            }
        }

        var exist = false
        for((i, row) in A.withIndex()){
            if(row.all {it == '?'}){
                if(exist){
                    A[i] = A[i-1]
                }
            }else{
                if(!exist){
                    exist = true
                    for(j in 0 until i)
                        A[j] = row
                }
            }
        }

        val res = A.map{it.joinToString("")}.joinToString("\n")
        println("Case #${it}:\n$res\n")
    }
}
