using System.Windows;

namespace PeopleViewer;

public partial class MainWindow : Window
{
    MainWindowViewModel viewModel = new();

    public MainWindow()
    {
        InitializeComponent();
        DataContext = viewModel;
    }

    private void RefreshButton_Click(object sender, RoutedEventArgs e)
    {
        viewModel.RefreshData();
    }
}
