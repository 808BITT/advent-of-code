object Part1 {
  def main(args: Array[String]): Unit = {
    import scala.io.Source
    import System.currentTimeMillis

    val startTimer = currentTimeMillis()
    val data = Source.fromFile("input.txt").getLines().toList

    var sum = 0
    for (line <- data) {
      var left, right: Option[Char] = None
      for (i <- line.indices) {
        if (line(i).isDigit) {
          if (left.isEmpty) {
            left = Some(line(i))
            right = Some(line(i))
          } else {
            right = Some(line(i))
          }
        }
      }
      sum += (left.get.asDigit * 10) + right.get.asDigit
    }
    println(sum)
    println(s"${(currentTimeMillis() - startTimer)} ms")
  }
}